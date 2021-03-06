/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// CommandIdentity represents the Identity of the object
var CommandIdentity = bambou.Identity{
	Name:     "command",
	Category: "commands",
}

// CommandsList represents a list of Commands
type CommandsList []*Command

// CommandsAncestor is the interface that an ancestor of a Command must implement.
// An Ancestor is defined as an entity that has Command as a descendant.
// An Ancestor can get a list of its child Commands, but not necessarily create one.
type CommandsAncestor interface {
	Commands(*bambou.FetchingInfo) (CommandsList, *bambou.Error)
}

// CommandsParent is the interface that a parent of a Command must implement.
// A Parent is defined as an entity that has Command as a child.
// A Parent is an Ancestor which can create a Command.
type CommandsParent interface {
	CommandsAncestor
	CreateCommand(*Command) *bambou.Error
}

// Command represents the model of a command
type Command struct {
	ID                  string `json:"ID,omitempty"`
	ParentID            string `json:"parentID,omitempty"`
	ParentType          string `json:"parentType,omitempty"`
	Owner               string `json:"owner,omitempty"`
	LastUpdatedBy       string `json:"lastUpdatedBy,omitempty"`
	DetailedStatus      string `json:"detailedStatus,omitempty"`
	DetailedStatusCode  int    `json:"detailedStatusCode,omitempty"`
	EntityScope         string `json:"entityScope,omitempty"`
	Command             string `json:"command,omitempty"`
	CommandInformation  string `json:"commandInformation,omitempty"`
	AssociatedParam     string `json:"associatedParam,omitempty"`
	AssociatedParamType string `json:"associatedParamType,omitempty"`
	Status              string `json:"status,omitempty"`
	FullCommand         string `json:"fullCommand,omitempty"`
	Summary             string `json:"summary,omitempty"`
	Override            string `json:"override,omitempty"`
	ExternalID          string `json:"externalID,omitempty"`
}

// NewCommand returns a new *Command
func NewCommand() *Command {

	return &Command{
		DetailedStatusCode: 0,
		Command:            "UNKNOWN",
		Status:             "UNKNOWN",
		Override:           "UNSPECIFIED",
	}
}

// Identity returns the Identity of the object.
func (o *Command) Identity() bambou.Identity {

	return CommandIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Command) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Command) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Command from the server
func (o *Command) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Command into the server
func (o *Command) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Command from the server
func (o *Command) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
