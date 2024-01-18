package pokedexapi

import (
	"encoding/json"
	"errors"
)


func getConfig(val []byte)(*Config, error){
	cfg := Config{}
	err := json.Unmarshal(val, &cfg)
	if err != nil {
		return &cfg, err
	}
	return &cfg, nil
}


func GetNextConfig(cfg *Config) (*Config, error){
   if cfg.Next == nil {
	  return cfg, errors.New("no next Page")
   }
   url := *cfg.Next
   val, err := getJson(url)
   if(err != nil){
	  return cfg, err
   }
   newCfg,err := getConfig(val)
   if(err != nil){
	return cfg, err
   }

   return  newCfg, nil
}

func GetPrevConfig(cfg *Config) (*Config, error){
	if cfg.Previous == nil {
	   return cfg, errors.New("no previous Page")
	}
	url := *cfg.Previous
	val, err := getJson(url)
	if(err != nil){
	   return cfg, err
	}
	newCfg,err := getConfig(val)
	if(err != nil){
	 return cfg, err
	}
 
	return  newCfg, nil
 }