package database

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/boltdb/bolt"
	"github.com/pspiagicw/goreland"
)

type Package struct {
	Name      string    `json:"name"`
	Checksum  string    `json:"checkSum"`
	Version   string    `json:"version"`
	Path      string    `json:"path"`
	Installed time.Time `json:"installed"`
}

type GoxDatabase struct {
	Packages map[string]Package `json:"packages"`
}

func getDatabasePath() string {
	location, exists := os.LookupEnv("XDG_DATA_HOME")
	if !exists {
		goreland.LogInfo("Not using $XDG_DATA_HOME, env variable not present")
		homedir, err := os.UserHomeDir()
		if err != nil {
			goreland.LogFatal("Error while getting $HOME directory: %v", err)
		}
		d := filepath.Join(homedir, ".local")
		d = filepath.Join(d, "share")
		d = filepath.Join(d, "gox")
		d = filepath.Join(d, "db")
		goreland.LogInfo("Using %s for database", d)
		return d
	}
	d := filepath.Join(location, "gox")
	d = filepath.Join(d, "db")
	goreland.LogInfo("Using %s for database", d)
	return d

}
func ParseDatabase() (*GoxDatabase, error) {
	path := getDatabasePath()

	database, err := readDatabase(path)

	if err != nil {
		return nil, err
	}

	return database, nil

}
func readPackages(b *bolt.Bucket, packages map[string]Package) {
	b.ForEach(func(k, v []byte) error {
		p := new(Package)
		err := json.Unmarshal(v, &p)
		if err != nil {
			return fmt.Errorf("Error unmarshalling struct :%v", err)
		}
		packages[string(k)] = *p
		return nil
	})
}

func readDatabase(path string) (*GoxDatabase, error) {

	gdb := new(GoxDatabase)
	gdb.Packages = map[string]Package{}

	db, err := bolt.Open(path, 0600, nil)
    if db != nil {
        defer db.Close()
    }

	if err != nil {
		return nil, fmt.Errorf("Error reading the database: %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("packages"))

		if err != nil || b == nil {
			return fmt.Errorf("Error opening bucket: %v", err)
		}

        readPackages(b, gdb.Packages)

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error reading database: %v", err)
	}

	return gdb, nil
}

func AddPackage(pack Package) error {
	path := getDatabasePath()

	db, err := bolt.Open(path, 0600, nil)
	defer db.Close()
	if err != nil {
		return fmt.Errorf("Error reading the database: %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("packages"))

		if err != nil || b == nil {
			return fmt.Errorf("Error creating bucket: %v", err)
		}
		contents, err := json.Marshal(pack)
		if err != nil {
			return fmt.Errorf("Error marshalling pacakge: %v", err)
		}
		err = b.Put([]byte(pack.Name), contents)
		if err != nil {
			return fmt.Errorf("Error inserting into database: %v", err)
		}
		return nil

	})
	return err
}
