package server

import (
  "github.com/jaredfolkins/badactor"
  "log"
)

func (ccc *ChaosDiscord) WhenJailed(a *badactor.Actor, r *badactor.Rule) error {
  // todo: mute user


  switch r.Name {
	case "spamBlock":
	  //idk
	case "spamKick":
		// todo: change code to acomedate Guild (discord channel identity)
	default:
		log.Println("Error: undefined kailor rule: " + r.Name)
	}
	return nil
}

func (ccc *ChaosDiscord) WhenTimeServed(a *badactor.Actor, r *badactor.Rule) error {
  // Do something here. Log, email, etc...
  //todo: unmite user
	return nil
}