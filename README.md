# golab3
Implementation des interfaces par des structures

1. interface
```type InterfaceName interface {
  Method1() 
  Method2() (StructureName, errors)
  Method3(params type) (StructureName, errors)
}
```

2. strcuture
```type StructureName struct {
  varName1 type
  varName2 type
}
```
3. structure implements interface

func (StrcutureName) Method1() {
  // instructions ...
}
func (StrcutureName) Method2() (StructureName, errors) {
  // instructions ...
  return structure, nil
}

func (StrcutureName) Method3(params type) (StructureName, errors) {
  // instructions ...
  return structure, nil
}
