# hold my files

**hmf** is a simple (proof-of-concept) tool to share files privately on the IPFS network.

> IPFS is a distributed system for storing and accessing files, websites, applications, and data.

* See more at [What is IPFS?](https://docs.ipfs.io/concepts/what-is-ipfs/)

**hmf** adds a small encryption layer on top of IPFS functionality to enable sharing files with anyone, anywhere on the planet, in a decentralized **and** private way.

Encryption is done using AES-256-GCM while the encryption key is a SHA-256 generated hash based on a PBKDF2 key derivation function applied to either a diceware-generated 10 word passphrase or a custom user-defined password.     

Install with go get:

```
go get github.com/miguelcnf/hmf
```
---

**Use cases**

* Requires a running IPFS local node - the easiest way to set it up is by installing [IPFS Desktop](https://docs.ipfs.io/install/ipfs-desktop/)

Share a file:

```
$ echo 'test' > test.txt 
$ hmf -share -file test.txt
Make sure to write down your 10 word generated passphrase: bloating licking echo strainer ought oil possibly coauthor litmus smudge
Here's your file identifier: QmVkjAfjqZMhktY1FpKQSfXD7LWQfVn6fZkoXt292QNF8a
Success!
```

Share a file with a custom password:

```
$ echo 'test' > test.txt 
$ hmf -share -file test.txt -password
Choose a passphrase:  
Confirm your passphrase: 
Here's your file identifier: QmVkjAfjqZMhktY1FpKQSfXD7LWQfVn6fZkoXt292QNF8a
Success!
```

Get (and decrypt) a shared file:

```
$ hmf -get -file test.txt
File identifier: QmVkjAfjqZMhktY1FpKQSfXD7LWQfVn6fZkoXt292QNF8a
Passphrase: [Generated or Custom passphrase]
Success! 
$ cat test.txt
test
```

Hold (and re-share) a shared file:

```
$ hmf -hold
File identifier: QmVkjAfjqZMhktY1FpKQSfXD7LWQfVn6fZkoXt292QNF8a
Success!
$ ipfs pin ls | grep QmVkjAfjqZMhktY1FpKQSfXD7LWQfVn6fZkoXt292QNF8a
QmVkjAfjqZMhktY1FpKQSfXD7LWQfVn6fZkoXt292QNF8a recursive
```

Read an encrypted shared file:

```
$ ipfs get QmVkjAfjqZMhktY1FpKQSfXD7LWQfVn6fZkoXt292QNF8a -o copy.txt
$ file copy.txt
copy.txt: data
$ hmf -read -file copy.txt 
Passphrase: [Generated or Custom passphrase]
Success!
$ file copy.txt
copy.txt: ASCII text
$ cat copy.txt
test
```




