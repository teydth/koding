package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/koding/kite"
	"github.com/koding/kite/config"
	"github.com/koding/kite/kontrol"
	"github.com/koding/kite/protocol"
	"github.com/koding/kite/testkeys"
	"github.com/koding/kite/testutil"

	"koding/db/models"
	"koding/db/mongodb/modelhelper"
	"koding/kites/kloud/keys"
	"koding/kites/kloud/kloud"
	"koding/kites/kloud/koding"
	"koding/kites/kloud/machinestate"
	"koding/kites/kloud/multiec2"
	kloudprotocol "koding/kites/kloud/protocol"
	"koding/kites/kloud/sshutil"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
)

var (
	kloudKite  *kite.Kite
	kld        *kloud.Kloud
	remote     *kite.Client
	conf       *config.Config
	provider   *koding.Provider
	machineId0 = "koding_id0"
)

const (
	etcdIp = "192.168.59.103:4001"
)

type args struct {
	MachineId string
}

func init() {
	conf = config.New()
	conf.Username = "testuser"
	conf.KontrolURL = "http://localhost:4099/kite"
	conf.KontrolKey = testkeys.Public
	conf.KontrolUser = "testuser"
	conf.KiteKey = testutil.NewKiteKey().Raw

	// Power up our own kontrol kite for self-contained tests
	kontrol.DefaultPort = 4099
	kntrl := kontrol.New(conf.Copy(), "0.1.0", testkeys.Public, testkeys.Private)
	kntrl.SetStorage(kontrol.NewPostgres(nil, kntrl.Kite.Log))

	go kntrl.Run()
	<-kntrl.Kite.ServerReadyNotify()

	// Power up kloud kite
	kloudKite = kite.New("kloud", "0.0.1")
	kloudKite.Config = conf.Copy()
	kloudKite.Config.Port = 4002
	kiteURL := &url.URL{Scheme: "http", Host: "localhost:4002", Path: "/kite"}
	_, err := kloudKite.Register(kiteURL)
	if err != nil {
		log.Fatal(err)
	}

	provider = newKodingProvider()

	// Add Kloud handlers
	kld := newKloud(provider)
	kloudKite.HandleFunc("build", kld.Build)
	kloudKite.HandleFunc("destroy", kld.Destroy)
	kloudKite.HandleFunc("start", kld.Start)
	kloudKite.HandleFunc("stop", kld.Stop)
	kloudKite.HandleFunc("reinit", kld.Reinit)
	kloudKite.HandleFunc("resize", kld.Resize)
	kloudKite.HandleFunc("event", kld.Event)

	go kloudKite.Run()
	<-kloudKite.ServerReadyNotify()

	user := kite.New("user", "0.0.1")
	user.Config = conf.Copy()

	kloudQuery := &protocol.KontrolQuery{
		Username:    "testuser",
		Environment: conf.Environment,
		Name:        "kloud",
	}
	kites, err := user.GetKites(kloudQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Get the caller
	remote = kites[0]
	if err := remote.Dial(); err != nil {
		log.Fatal(err)
	}
}

// Main VM action tests (build, start, stop, destroy, resize, reinit)
func TestPing(t *testing.T) {
	_, err := remote.Tell("kite.ping")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuild(t *testing.T) {
	userData, err := createUser()
	if err != nil {
		t.Fatal(err)
	}

	if err := build(userData.MachineId); err != nil {
		t.Error(err)
	}

	// now try to ssh into the machine with temporary private key we created in
	// the beginning
	if err := checkSSHKey(userData.MachineId, userData.PrivateKey); err != nil {
		t.Error(err)
	}
}

func TestInvalidMethodsOnRunning(t *testing.T) {
	t.Log("Running invalid methods on a running VM.")
	if err := build(machineId0); err == nil {
		t.Error("`build` method can not be called on `running` machines.")
	}
}

func TestStop(t *testing.T) {
	if err := stop(machineId0); err != nil {
		t.Error(err)
	}
}

func TestInvalidMethodsOnStopped(t *testing.T) {
	t.Log("Running invalid methods on a stopped VM.")
	// run the tests now.
	if err := build(machineId0); err == nil {
		t.Error("`build` method can not be called on `stopped` machines.")
	}

	if err := stop(machineId0); err == nil {
		t.Error("`stop` method can not be called on `stopped` machines.")
	}
}

func TestStart(t *testing.T) {
	if err := start(machineId0); err != nil {
		t.Error(err)
	}
}

func TestResize(t *testing.T) {
	storageWant := 5
	m := GetMachineData(machineId0)
	m.Builder["storage_size"] = storageWant
	SetMachineData(machineId0, m)

	if err := resize(machineId0); err != nil {
		t.Error(err)
	}

	storageGot, err := getAmazonStorageSize(machineId0)
	if err != nil {
		t.Error(err)
	}

	if storageGot != storageWant {
		t.Errorf("Resizing completed but storage sizes do not match. Want: %dGB, Got: %dGB",
			storageWant,
			storageGot,
		)
	}
}

func TestReinit(t *testing.T) {
	if err := reinit(machineId0); err != nil {
		t.Error(err)
	}
}

func TestDestroy(t *testing.T) {
	if err := destroy(machineId0); err != nil {
		t.Error(err)
	}
}

func TestInvalidMethodsOnTerminated(t *testing.T) {
	t.Log("Running invalid methods on a terminated VM.")
	if err := stop(machineId0); err == nil {
		t.Error("`stop` method can not be called on `terminated` machines.")
	}

	if err := start(machineId0); err == nil {
		t.Error("`start` method can not be called on `terminated` machines.")
	}

	if err := destroy(machineId0); err == nil {
		t.Error("`destroy` method can not be called on `terminated` machines.")
	}

	if err := resize(machineId0); err == nil {
		t.Error("`resize` method can not be called on `terminated` machines.")
	}

	if err := reinit(machineId0); err == nil {
		t.Error("`reinit` method can not be called on `terminated` machines.")
	}
}

type singleUser struct {
	MachineId  string
	PrivateKey string
	PublicKey  string
}

// createUser creates a test user in jUsers and a single jMachine document.
func createUser() (*singleUser, error) {
	privateKey, publicKey, err := sshutil.TemporaryKey()
	if err != nil {
		return nil, err
	}

	username := "testuser"

	userId := bson.NewObjectId()
	user := &models.User{
		ObjectId:      userId,
		Email:         "testuser@testuser.com",
		LastLoginDate: time.Now().UTC(),
		RegisteredAt:  time.Now().UTC(),
		Name:          username, // bson equivelant is username
		Password:      "somerandomnumbers",
		Status:        "confirmed",
		SshKeys: []struct {
			Title string `bson:"title"`
			Key   string `bson:"key"`
		}{
			{Key: publicKey},
		},
	}

	if err := provider.Session.Run("jUsers", func(c *mgo.Collection) error {
		return c.Insert(&user)
	}); err != nil {
		return nil, err
	}

	// later we can add more users with "Owner:false" to test sharing capabilities
	users := []models.Permissions{
		{Id: userId, Sudo: true, Owner: true},
	}

	machineId := bson.NewObjectId()
	machine := &koding.MachineDocument{
		Id:         machineId,
		Label:      "",
		Domain:     username + ".dev.koding.io",
		Credential: username,
		Provider:   "koding",
		CreatedAt:  time.Now().UTC(),
		Meta: bson.M{
			"region":        "eu-west-1",
			"instance_type": "t2.micro",
			"storage_size":  3,
			"alwaysOn":      false,
		},
		Users:  users,
		Groups: make([]models.Permissions, 0),
	}
	machine.Assignee.InProgress = false
	machine.Assignee.AssignedAt = time.Now().UTC()
	machine.Status.State = machinestate.NotInitialized.String()
	machine.Status.ModifiedAt = time.Now().UTC()

	if err := provider.Session.Run("jMachines", func(c *mgo.Collection) error {
		return c.Insert(&machine)
	}); err != nil {
		return nil, err
	}

	return &singleUser{
		MachineId:  machineId.Hex(),
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func build(id string) error {
	buildArgs := &args{
		MachineId: id,
	}

	resp, err := remote.Tell("build", buildArgs)
	if err != nil {
		return err
	}

	var result kloud.ControlResult
	err = resp.Unmarshal(&result)
	if err != nil {
		return err
	}

	eArgs := kloud.EventArgs([]kloud.EventArg{
		kloud.EventArg{
			EventId: buildArgs.MachineId,
			Type:    "build",
		},
	})

	return listenEvent(eArgs, machinestate.Running)

}

func checkSSHKey(id, privateKey string) error {
	// now try to ssh into the machine with temporary private key we created in
	// the beginning
	machine, err := provider.Get(id)
	if err != nil {
		return err
	}

	sshConfig, err := sshutil.SshConfig("root", privateKey)
	if err != nil {
		return err
	}

	log.Printf("Connecting to machine with ip '%s' via ssh\n", machine.IpAddress)
	sshClient, err := sshutil.ConnectSSH(machine.IpAddress+":22", sshConfig)
	if err != nil {
		return err
	}

	output, err := sshClient.StartCommand("whoami")
	if err != nil {
		return err
	}

	if strings.TrimSpace(string(output)) != "root" {
		return fmt.Errorf("Whoami result should be root, got: %s", string(output))
	}

	return nil
}

func destroy(id string) error {
	destroyArgs := &args{
		MachineId: id,
	}

	resp, err := remote.Tell("destroy", destroyArgs)
	if err != nil {
		return err
	}

	var result kloud.ControlResult
	err = resp.Unmarshal(&result)
	if err != nil {
		return err
	}
	eArgs := kloud.EventArgs([]kloud.EventArg{
		kloud.EventArg{
			EventId: destroyArgs.MachineId,
			Type:    "destroy",
		},
	})

	if err := listenEvent(eArgs, machinestate.Terminated); err != nil {
		return err
	}

	return nil
}

func start(id string) error {
	startArgs := &args{
		MachineId: id,
	}

	resp, err := remote.Tell("start", startArgs)
	if err != nil {
		return err
	}

	var result kloud.ControlResult
	err = resp.Unmarshal(&result)
	if err != nil {
		return err
	}

	eArgs := kloud.EventArgs([]kloud.EventArg{
		kloud.EventArg{
			EventId: startArgs.MachineId,
			Type:    "start",
		},
	})

	if err := listenEvent(eArgs, machinestate.Running); err != nil {
		return err
	}

	return nil
}

func stop(id string) error {
	stopArgs := &args{
		MachineId: id,
	}

	resp, err := remote.Tell("stop", stopArgs)
	if err != nil {
		return err
	}

	var result kloud.ControlResult
	err = resp.Unmarshal(&result)
	if err != nil {
		return err
	}

	eArgs := kloud.EventArgs([]kloud.EventArg{
		kloud.EventArg{
			EventId: stopArgs.MachineId,
			Type:    "stop",
		},
	})

	if err := listenEvent(eArgs, machinestate.Stopped); err != nil {
		return err
	}

	return nil
}

func reinit(id string) error {
	reinitArgs := &args{
		MachineId: id,
	}

	resp, err := remote.Tell("reinit", reinitArgs)
	if err != nil {
		return err
	}

	var result kloud.ControlResult
	err = resp.Unmarshal(&result)
	if err != nil {
		return err
	}

	eArgs := kloud.EventArgs([]kloud.EventArg{
		kloud.EventArg{
			EventId: reinitArgs.MachineId,
			Type:    "reinit",
		},
	})

	if err := listenEvent(eArgs, machinestate.Running); err != nil {
		return err
	}

	return nil
}

func resize(id string) error {
	resizeArgs := &args{
		MachineId: id,
	}

	resp, err := remote.Tell("resize", resizeArgs)
	if err != nil {
		return err
	}

	var result kloud.ControlResult
	err = resp.Unmarshal(&result)
	if err != nil {
		return err
	}

	eArgs := kloud.EventArgs([]kloud.EventArg{
		kloud.EventArg{
			EventId: resizeArgs.MachineId,
			Type:    "resize",
		},
	})

	if err := listenEvent(eArgs, machinestate.Running); err != nil {
		return err
	}

	return nil
}

// listenEvent calls the event method of kloud with the given arguments until
// the desiredState is received. It times out if the desired state is not
// reached in 10 miunuts.
func listenEvent(args kloud.EventArgs, desiredState machinestate.State) error {
	tryUntil := time.Now().Add(time.Minute * 10)
	for {
		resp, err := remote.Tell("event", args)
		if err != nil {
			return err
		}

		var events []kloud.EventResponse
		if err := resp.Unmarshal(&events); err != nil {
			return err
		}

		e := events[0]
		if e.Error != nil {
			return e.Error
		}

		event := e.Event

		if event.Status == desiredState {
			return nil
		}

		if time.Now().After(tryUntil) {
			return fmt.Errorf("Timeout while waiting for state %s", desiredState)
		}

		time.Sleep(2 * time.Second)
		continue // still pending
	}

	return nil
}

func newKodingProvider() *koding.Provider {
	auth := aws.Auth{
		AccessKey: "AKIAJFKDHRJ7Q5G4MOUQ",
		SecretKey: "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",
	}

	mongoURL := os.Getenv("KLOUD_MONGODB_URL")
	if mongoURL == "" {
		panic("KLOUD_MONGODB_URL is not set")
	}

	modelhelper.Initialize(mongoURL)
	db := modelhelper.Mongo
	domainStorage := koding.NewDomainStorage(db)

	testChecker := &TestChecker{}

	return &koding.Provider{
		Session:           db,
		Kite:              kloudKite,
		Log:               newLogger("koding", false),
		KontrolURL:        conf.KontrolURL,
		KontrolPrivateKey: testkeys.Private,
		KontrolPublicKey:  testkeys.Public,
		Test:              true,
		EC2Clients: multiec2.New(auth, []string{
			"us-east-1",
			"ap-southeast-1",
			"us-west-2",
			"eu-west-1",
		}),
		DNS:           koding.NewDNSClient("dev.koding.io", auth),
		DomainStorage: domainStorage,
		Bucket:        koding.NewBucket("koding-klient", "development/latest", auth),
		KeyName:       keys.DeployKeyName,
		PublicKey:     keys.DeployPublicKey,
		PrivateKey:    keys.DeployPrivateKey,
		PlanChecker: func(_ *kloudprotocol.Machine) (koding.Checker, error) {
			return testChecker, nil
		},
	}
}

func newKloud(p *koding.Provider) *kloud.Kloud {
	kld := kloud.New()
	kld.Log = newLogger("kloud", false)
	kld.Locker = p
	kld.Storage = p
	kld.DomainStorage = p.DomainStorage
	kld.Domainer = p.DNS
	kld.Debug = true
	kld.AddProvider("koding", p)
	return kld
}

func getAmazonStorageSize(machineId string) (int, error) {
	m := GetMachineData(machineId0)

	a, err := provider.NewClient(m)
	if err != nil {
		return 0, err
	}

	instance, err := a.Instance(a.Id())
	if err != nil {
		return 0, err
	}

	if len(instance.BlockDevices) == 0 {
		return 0, fmt.Errorf("fatal error: no block device available")
	}

	// we need in a lot of placages!
	oldVolumeId := instance.BlockDevices[0].VolumeId

	oldVolResp, err := a.Client.Volumes([]string{oldVolumeId}, ec2.NewFilter())
	if err != nil {
		return 0, err
	}

	volSize := oldVolResp.Volumes[0].Size
	currentSize, err := strconv.Atoi(volSize)
	if err != nil {
		return 0, err
	}

	return currentSize, nil
}
