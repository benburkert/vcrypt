package test

var (
	// DiamondPlanConfig is a diamond shaped password plan
	DiamondPlanConfig = []byte(`
# [secretbox "top"] -> [password "top password"]
#                   |
#                   -> [mux] -> [secretbox "left"]  -> [password "left password"]
#                            |                      |
#                            |                      -> [demux] -> [secretbox "bottom"] -> [password "bottom password"]
#                            |                      |                                  |
#                            -> [secretbox "right"] -> [password "right password"]     -> [material "bottom material"]

comment = Diamond shaped plan
root = top

[secretbox "top"]
comment = step 3
edge = top password
edge = mux

[mux]
edge = left
edge = right

[secretbox "left"]
comment = step 2a
edge = left password
edge = demux

[secretbox "right"]
comment = step 2b
edge = right password
edge = demux

[demux]
edge = bottom

[secretbox "bottom"]
comment = step 1
edge = bottom password
edge = bottom material

[password "top password"]
comment = step 3 password

[password "left password"]
comment = step 2a password

[password "right password"]
comment = step 2b password

[password "bottom password"]
comment = step 1 password

[material "bottom material"]
`)

	// TwoManPlanConfig represents a two-man encryption scheme plan
	TwoManPlanConfig = []byte(`
# [secretbox "master key"] -> [secretbox "op 1 key"] -> [password "op 1 password"]
#                          |                         |
#                          |                         -> [material "op 1 material"]
#                          |
#                          -> [secretbox "op 2 key"] -> [password "op 2 password"]
#                                                    |
#                                                    -> [material "op 2 material"]

comment = Two-man rule plan
root = master key

[secretbox "master key"]
edge = op 1 key
edge = op 2 key

[secretbox "op 1 key"]
comment = operator 1 key
edge = op 1 password
edge = op 1 material

[secretbox "op 2 key"]
comment = operator 2 key
edge = op 2 password
edge = op 2 material

[password "op 1 password"]
comment = op 1 secret

[password "op 2 password"]
comment = op 2 secret

[material "op 1 material"]

[material "op 2 material"]
`)

	// TwoPartyPlanConfig is a request-approve-unlock encryption plan
	TwoPartyPlanConfig = []byte(`
# [secretbox "step 3"] -> [password "step 3 password"]
#                      |
#                      -> [secretbox "step 2"] -> [password "step 2 password"]
#                                              |
#                                              -> [secretbox "step 1"] -> [password "step 1 password"]
#                                                                      |
#                                                                      -> [material]

comment = Two-party 3 step plan
root = step 3

[secretbox "step 3"]
edge = step 3 password
edge = step 2

[secretbox "step 2"]
edge = step 2 password
edge = step 1

[secretbox "step 1"]
edge = step 1 password
edge = material

[password "step 3 password"]
comment = party 1 password 2

[password "step 2 password"]
comment = party 2 password

[password "step 1 password"]
comment = party 1 password 1

[material]
`)

	// DNSSecConfig is a multi-party m-of-n OpenPGP encryption plan inspired by
	// the DNSSEC root key.
	DNSSecConfig = []byte(`
# [sss "five-of-seven"] -> [openpgp "alice@example.com"] -> [openpgp-key "alice@example.com"]
#                       |                                |
#                       |                                -> [material]
#                       |
#                       -> [openpgp "bob@example.com"] -> [openpgp-key "bob@example.com"]
#                       |                              |
#                       |                              -> [material]
#                       |
#                       -> [openpgp "claire@example.com"] -> [openpgp-key "claire@example.com"]
#                       |                                 |
#                       |                                 -> [material]
#                       |
#                       -> [openpgp "david@example.com"] -> [openpgp-key "david@example.com"]
#                       |                                |
#                       |                                -> [material]
#                       |
#                       -> [openpgp "emily@example.com"] -> [openpgp-key "emily@example.com"]
#                       |                                |
#                       |                                -> [material]
#                       |
#                       -> [openpgp "frank@example.com"] -> [openpgp-key "frank@example.com"]
#                       |                                |
#                       |                                -> [material]
#                       |
#                       -> [openpgp "gloria@example.com"] -> [openpgp-key "gloria@example.com"]
#                                                         |
#                                                         -> [material]

comment = DNSSEC Root Key
root = five-of-seven

[sss "five-of-seven"]
max-shares = 7
required-shares = 5
edge = alice@example.com
edge = bob@example.com
edge = claire@example.com
edge = david@example.com
edge = emily@example.com
edge = frank@example.com
edge = gloria@example.com

[openpgp "alice@example.com"]
publickey = "` + OpenPGPKeys["alice"].Public + `"
edge = alice material
edge = ` + OpenPGPKeys["alice"].KeyID + `

[openpgp "bob@example.com"]
publickey = "` + OpenPGPKeys["bob"].Public + `"
edge = bob material
edge = ` + OpenPGPKeys["bob"].KeyID + `

[openpgp "claire@example.com"]
publickey = "` + OpenPGPKeys["claire"].Public + `"
edge = claire material
edge = ` + OpenPGPKeys["claire"].KeyID + `

[openpgp "david@example.com"]
publickey = "` + OpenPGPKeys["david"].Public + `"
edge = david material
edge = ` + OpenPGPKeys["david"].KeyID + `

[openpgp "emily@example.com"]
publickey = "` + OpenPGPKeys["emily"].Public + `"
edge = emily material
edge = ` + OpenPGPKeys["emily"].KeyID + `

[openpgp "frank@example.com"]
publickey = "` + OpenPGPKeys["frank"].Public + `"
edge = frank material
edge = ` + OpenPGPKeys["frank"].KeyID + `

[openpgp "gloria@example.com"]
publickey = "` + OpenPGPKeys["gloria"].Public + `"
edge = gloria material
edge = ` + OpenPGPKeys["gloria"].KeyID + `

[openpgp-key "` + OpenPGPKeys["alice"].KeyID + `"]
comment = ` + OpenPGPKeys["alice"].KeyID + `
keyid = ` + OpenPGPKeys["alice"].KeyID + `

[openpgp-key "` + OpenPGPKeys["bob"].KeyID + `"]
comment = ` + OpenPGPKeys["bob"].KeyID + `
keyid = ` + OpenPGPKeys["bob"].KeyID + `

[openpgp-key "` + OpenPGPKeys["claire"].KeyID + `"]
comment = ` + OpenPGPKeys["claire"].KeyID + `
keyid = ` + OpenPGPKeys["claire"].KeyID + `

[openpgp-key "` + OpenPGPKeys["david"].KeyID + `"]
comment = ` + OpenPGPKeys["david"].KeyID + `
keyid = ` + OpenPGPKeys["david"].KeyID + `

[openpgp-key "` + OpenPGPKeys["emily"].KeyID + `"]
comment = ` + OpenPGPKeys["emily"].KeyID + `
keyid = ` + OpenPGPKeys["emily"].KeyID + `

[openpgp-key "` + OpenPGPKeys["frank"].KeyID + `"]
comment = ` + OpenPGPKeys["frank"].KeyID + `
keyid = ` + OpenPGPKeys["frank"].KeyID + `

[openpgp-key "` + OpenPGPKeys["gloria"].KeyID + `"]
comment = ` + OpenPGPKeys["gloria"].KeyID + `
keyid = ` + OpenPGPKeys["gloria"].KeyID + `

[material "alice material"]

[material "bob material"]

[material "claire material"]

[material "david material"]

[material "emily material"]

[material "frank material"]

[material "gloria material"]

`)
)
