# Ukko

The standard Merlin command wrapper

`Usage: ukko [-rReEmh] [FILE]...`

`-r`: Enter raw mode, parse stdin and allow interupts

`-R`: Like `-r` but with the standard `SPELLBOOK` loaded

`-e`: Evaluate a single file

`-E`: `-e` with the `SPELLBOOk` loaded

`-m`: Don't load `.manifest.mn`

`-h`: Display help

Enviroment variables: `SPELLBOOK` (location for your default spellbook)

Dynamic configurations: `ukko` will load any file named `.manifest.mn` in the current directory. This can be disabled with the `-m` flag.
