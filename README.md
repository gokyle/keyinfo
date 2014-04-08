keyinfo: display information about private keys

Usage: keyinfo [-v] files...

The `-v` option will dump the `D` and `N` values of RSA private keys.

Example:

```
keyinfo ec_private.pem
[+] Type: EC
[+] Curve: P256
```
With `-v`:

```
$ keyinfo -v private.key
[+] Type: RSA
[+] Size: 512
D: 8764fc094381bbd99e9e5b25aeb34aff6873400134f2b9c697051283095ace1b332f7832b9d8abcb6926b3f674ca62c194b7ee35a98882a621451b3fcd1705d1
E: b45c23d514571f9a3bac566604b02d6364c8bb2e6e25854f261916b2b501c188a9d512908cfbf972ae50c5e9ab0307a17d3ac1eb4a9ffa9d6422f0a1a03c69df
```
