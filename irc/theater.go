// Copyright (c) 2012-2014 Jeremy Latt
// Copyright (c) 2014-2015 Edmund Huber
// released under the MIT license

package irc

type TheaterClient Name

func (c TheaterClient) Id() Name {
	return Name(c)
}

func (c TheaterClient) Nick() Name {
	return Name(c)
}

/*
func (m *TheaterIdentifyCommand) LoadPassword(s *Server) {
	m.hash = s.theaters[m.channel]
}

	if upperSubCmd := strings.ToUpper(args[0]); upperSubCmd == "IDENTIFY" && len(args) == 3 {
		return &TheaterIdentifyCommand{
			channel:     NewName(args[1]),
			PassCommand: PassCommand{password: []byte(args[2])},
		}, nil
	} else if upperSubCmd == "PRIVMSG" && len(args) == 4 {
		return &TheaterPrivMsgCommand{
			channel: NewName(args[1]),
			asNick:  NewName(args[2]),
			message: NewText(args[3]),
		}, nil
	} else if upperSubCmd == "ACTION" && len(args) == 4 {
		return &TheaterActionCommand{
			channel: NewName(args[1]),
			asNick:  NewName(args[2]),
			action:  NewCTCPText(args[3]),
		}, nil
	} else {
		return nil, ErrParseCommand
	}
func (m *TheaterIdentifyCommand) HandleServer(s *Server) {
	client := m.Client()
	if !m.channel.IsChannel() {
		client.ErrNoSuchChannel(m.channel)
		return
	}

	channel := s.channels.Get(m.channel)
	if channel == nil {
		client.ErrNoSuchChannel(m.channel)
		return
	}

	if (m.hash == nil) || (m.err != nil) {
		client.ErrPasswdMismatch()
		return
	}

	if channel.members.AnyHasMode(Theater) {
		client.Reply(RplNotice(s, client, "someone else is +T in this channel"))
		return
	}

	channel.members[client][Theater] = true
}

type TheaterPrivMsgCommand struct {
	BaseCommand
	channel Name
	asNick  Name
	message Text
}

func (m *TheaterPrivMsgCommand) HandleServer(s *Server) {
	client := m.Client()

	if !m.channel.IsChannel() {
		client.ErrNoSuchChannel(m.channel)
		return
	}

	channel := s.channels.Get(m.channel)
	if channel == nil {
		client.ErrNoSuchChannel(m.channel)
		return
	}

	if !channel.members.HasMode(client, Theater) {
		client.Reply(RplNotice(s, client, "you are not +T"))
		return
	}

	reply := RplPrivMsg(TheaterClient(m.asNick), channel, m.message)
	for member := range channel.members {
		member.Reply(reply)
	}
}

type TheaterActionCommand struct {
	BaseCommand
	channel Name
	asNick  Name
	action  CTCPText
}

func (m *TheaterActionCommand) HandleServer(s *Server) {
	client := m.Client()

	if !m.channel.IsChannel() {
		client.ErrNoSuchChannel(m.channel)
		return
	}

	channel := s.channels.Get(m.channel)
	if channel == nil {
		client.ErrNoSuchChannel(m.channel)
		return
	}

	if !channel.members.HasMode(client, Theater) {
		client.Reply(RplNotice(s, client, "you are not +T"))
		return
	}

	reply := RplCTCPAction(TheaterClient(m.asNick), channel, m.action)
	for member := range channel.members {
		member.Reply(reply)
	}
}
*/