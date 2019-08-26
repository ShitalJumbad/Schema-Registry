
## Package goavro
Package goavro is a library that encodes and decodes Avro data.

type Codec
1) func NewCodec(schemaSpecification string) (*Codec, error)
  NewCodec returns a Codec used to translate between a byte slice of either binary or textual Avro data and native Go data.
  Internally a `Codec` is merely a named tuple of four function pointers, and maintains no runtime state that is mutated after instantiation. 
  In other words, `Codec`s may be safely used by many go routines simultaneously, as your program requires
  
2) func (c *Codec) NativeFromTextual(buf []byte) (interface{}, []byte, error)
  NativeFromTextual converts Avro data in JSON text format from the provided byte slice to Go native data types in accordance with the Avro schema supplied when creating the Codec.
  it returns the decoded datum, along with a new byte slice with the decoded bytes consumed
  
3) func (c *Codec) BinaryFromNative(buf []byte, datum interface{}) ([]byte, error)
  BinaryFromNative appends the binary encoded byte slice representation of the provided native datum value to the provided byte slice in accordance with the 
  Avro schema supplied when creating the Codec.
  
  





References:

https://godoc.org/github.com/linkedin/goavro
