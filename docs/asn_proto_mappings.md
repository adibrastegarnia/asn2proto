# ASN.1 to Proto Mappings


## Trivial Conversions

### ENUMERATED

`ENUMERATED` type in ASN.1 can be converted to `enum` directly in protobuf.
For example, 


```asn.1
Fuel ::= ENUMERATED {
    solid(0),
    liquid(1),
    gas(2),
    hybrid(3)
}
```

to

```protobuf
enum Fuel {
        solid = 0;
        liquid = 1;
        gas = 2;
        hybrid = 3;
}
```



### CHOICE OF

`CHOICE OF` type in ASN.1 can be converted to `oneof` in protobuf

For example 

```asn.1
speed CHOICE {
        mph INTEGER,
        kmph INTEGER
    },
```

to 

```protobuf
oneof speed  { 
          int32 mph = 4; 
          int32 kmph = 5;
        }  
```

### SEQUENCE 

`SEQUENCE` type in ASN.1 can be mapped to `message` in protobuf

```asn.1
RepeatedExample ::= SEQUENCE {
    repeated_field SEQUENCE OF INTEGER
}
```


```protobuf 
message RepeatedExample {
   repeated int32 repeated_field = 1;

}
```


### SEQUENCE OF

`SEQUENCE OF` type in ASN.1 to `repeated` in protobuf

```asn.1
RepeatedExample ::= SEQUENCE {
    repeated_field SEQUENCE OF INTEGER
}
```

```protobuf
message RepeatedExample {
   repeated int32 repeated_field = 1;

}
```


### SEQUENCE OF SEQUENCE

`SEQUENCE OF SEQUENCE` type in ASN.1 can be mapped to `map` in protobuf 

```asn.1
MapExample ::= SEQUENCE {
    my_map SEQUENCE OF SEQUENCE {
        key INTEGER,
        value UTF8String
    }
}
```


```protobuf
message MapExample {
  map<int32, string> my_map = 4;

}
```



## Non Trivial Conversions