package main

import (
    "os"
    "encoding/json"
)

// Config is the global config struct for site.
type Config struct {
    // BasePath is the root of our site's URL on server.
    BasePath string
}

// Load loads the global config from a file.
func (cfg *Config) Load(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()

    dec := json.NewDecoder(f)
    if err := dec.Decode(cfg); err != nil {
        return err
    }

    return nil
}

