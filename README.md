# tcp-test
Custom binary tcp protocol

## Protocol to implement:

```
<msg_id:uint32><seq:uint32><data_len:uint32><data_json:[]byte><data_crc32:uint32>

msg_id - command code
seq - query id
data_len - length of data
data_json - json data
data_crc32 - hash sum of data
```

First we need to encode data and write it to tcp connection.
After we can write to tcp conn we need to read from tcp conn.
