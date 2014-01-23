package capn

const Package = uint64(0xbea97f1023792be0)
const Import = uint64(0xe130b601260e44b5)
const Tag = uint64(0xa574b41924caefc7)
const Notag = uint64(0xc8768679ec52e012)
const Jsonfmt = uint64(0xe5c80e8fc92442ef)

type Jsonfmts Struct

func NewJsonfmts(s *Segment) Jsonfmts      { return Jsonfmts(s.NewStruct(0, 2)) }
func NewRootJsonfmts(s *Segment) Jsonfmts  { return Jsonfmts(s.NewRootStruct(0, 2)) }
func ReadRootJsonfmts(s *Segment) Jsonfmts { return Jsonfmts(s.Root(0).ToStruct()) }
func (s Jsonfmts) Pkg() string               { return Struct(s).GetObject(0).ToText() }
func (s Jsonfmts) SetPkg(v string)           { Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Jsonfmts) Fmt() string               { return Struct(s).GetObject(1).ToText() }
func (s Jsonfmts) SetFmt(v string)           { Struct(s).SetObject(1, s.Segment.NewText(v)) }
