package pdmq

import (
	"encoding/binary"
	"io"
	"regexp"
)

// MagicV1 is the initial identifier sent when connecting for V1 clients
var MagicV1 = []byte("V1")

// MagicV2 is the initial identifier sent when connecting for V2 clients
var MagicV2 = []byte("V2")

// frame types
const (
	FrameTypeResponse int32 = 0
	FrameTypeError    int32 = 1
	FrameTypeMessage  int32 = 2
)

var validTopicChannelNameRegex = regexp.MustCompile(`^[\.a-zA-Z0-9_-]+(#ephemeral)?$`)

// IsValidTopicName checks a topic name for correctness
func IsValidTopicName(name string) bool {
	return isValidName(name)
}

// IsValidChannelName checks a channel name for correctness
func IsValidChannelName(name string) bool {
	return isValidName(name)
}

func isValidName(name string) bool {
	if len(name) > 64 || len(name) < 1 {
		return false
	}
	return validTopicChannelNameRegex.MatchString(name)
}

// ReadResponse is a client-side utility function to read from the supplied Reader
// according to the PDMQ protocol spec:
//
//    [x][x][x][x][x][x][x][x]...
//    |  (int32) || (binary)
//    |  4-byte  || N-byte
//    ------------------------...
//        size       data
func ReadResponse(r io.Reader) ([]byte, error) {
	var msgSize int32

	// message size
	err := binary.Read(r, binary.BigEndian, &msgSize)
	if err != nil {
		return nil, err
	}

	// message binary data
	buf := make([]byte, msgSize)
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// UnpackResponse is a client-side utility function that unpacks serialized data
// according to PDMQ protocol spec:
//
//    [x][x][x][x][x][x][x][x]...
//    |  (int32) || (binary)
//    |  4-byte  || N-byte
//    ------------------------...
//      frame ID     data
//
// Returns a triplicate of: frame type, data ([]byte), error
func UnpackResponse(response []byte) (int32, []byte, error) {
	/*	if len(response) < 4 {
		return -1, nil, errors.New("length of response is too small")
	}*/

	return int32(binary.BigEndian.Uint32(response)), response, nil
}

// ReadUnpackedResponse reads and parses data from the underlying
// TCP connection according to the PDMQ TCP protocol spec and
// returns the frameType, data or error
func ReadUnpackedResponse(r io.Reader) (int32, []byte, error) {
	buf := make([]byte, 1024)
	bufLen, err := r.Read(buf)
	if err != nil {
		return -1, nil, err
	}
	return UnpackResponse(buf[:bufLen])
}

//这里是接收pdmqd 返回的所有信息，因为上游统一返回的是 message结构的[]byte数组，可以直接返回
//这里跟nsq 不一样
func ReadFullResponse(r io.Reader) ([]byte, error) {
	buf := make([]byte, 1024)
	bufLen, err := r.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:bufLen], nil
}
