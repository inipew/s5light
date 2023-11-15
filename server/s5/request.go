package s5

import (
	"io"
	"log"
	"net"

	"github.com/txthinking/socks5"
)

// Connect remote conn which u want to connect with your dialer
// Error or OK both replied.
func Connect(w io.Writer, r *socks5.Request, outAddress string) (*net.TCPConn, error) {
	if socks5.Debug {
		log.Printf("Handling Request: %s", r.Address())
	}

	tcpaddr, err := net.ResolveTCPAddr("tcp", r.Address())
	if err != nil {
		log.Printf("Failed to resolve TCP address: %v", err)
		replyWithError(w, r.Atyp, socks5.RepHostUnreachable)
		return nil, err
	}

	var localAddr *net.TCPAddr
	if outAddress != "" {
		localAddr, _ = net.ResolveTCPAddr("tcp", outAddress+":0")
	}

	rc, err := dialTCP(localAddr, tcpaddr)
	if err != nil {
		log.Printf("Failed to dial TCP: %v", err)
		replyWithError(w, r.Atyp, socks5.RepHostUnreachable)
		return nil, err
	}

	log.Printf("Successfully connected to: %s", rc.RemoteAddr())
	replyWithSuccess(w, rc.LocalAddr().String())
	return rc, nil
}

func dialTCP(localAddr, tcpaddr *net.TCPAddr) (*net.TCPConn, error) {
	var tmp *net.TCPConn
	var err error

	if localAddr != nil {
		tmp, err = socks5.Dial.DialTCP("tcp", localAddr, tcpaddr)
	} else {
		tmp, err = socks5.Dial.DialTCP("tcp", nil, tcpaddr)
	}

	return tmp, err
}

func replyWithError(w io.Writer, atyp uint8, rep uint8) {
	var p *socks5.Reply
	if atyp == socks5.ATYPIPv4 || atyp == socks5.ATYPDomain {
		p = socks5.NewReply(rep, socks5.ATYPIPv4, []byte{0x00, 0x00, 0x00, 0x00}, []byte{0x00, 0x00})
	} else {
		p = socks5.NewReply(rep, socks5.ATYPIPv6, []byte(net.IPv6zero), []byte{0x00, 0x00})
	}

	if _, err := p.WriteTo(w); err != nil {
		log.Println("Error writing reply:", err)
	}
}

func replyWithSuccess(w io.Writer, localAddr string) {
	a, addr, port, err := socks5.ParseAddress(localAddr)
	if err != nil {
		log.Printf("Failed to parse local address: %v", err)
		replyWithError(w, socks5.ATYPIPv4, socks5.RepHostUnreachable)
		return
	}

	p := socks5.NewReply(socks5.RepSuccess, a, addr, port)
	if _, err := p.WriteTo(w); err != nil {
		log.Println("Error writing reply:", err)
	}
}
