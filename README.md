# API Dumpster Fire

This is an example of composable APIs using [github.com/spudtrooper/minimalcli](https://github.com/spudtrooper/minimalcli) to create a little server hosting a few APIs:

https://api-dumpster-fire.herokuapp.com/

It currently contains unofficial APIs from opentable and resy:

| Site                                      | Individual Server                                                                         | Individual Project                                                               |
| ----------------------------------------- | ----------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| [opensecrets.org](http://opensecrets.org) | --                                                                                        | [github.com/spudtrooper/opensecrets](https://github.com/spudtrooper/opensecrets) |
| [opentable.com](http://opentable.com)     | [unofficial-opentable-api.herokuapp.com](https://unofficial-opentable-api.herokuapp.com/) | [github.com/spudtrooper/opentable](https://github.com/spudtrooper/opentable)     |
| [resy.com](http://resy.com)               | [unofficial-resy-api.herokuapp.com](https://unofficial-resy-api.herokuapp.com/)           | [github.com/spudtrooper/resy](https://github.com/spudtrooper/resy)               |
| [lyft.com](http://lyft.com)               | --                                                                                        | [github.com/spudtrooper/lyft](https://github.com/spudtrooper/lyft)               |

## Usage

### Deployed frontend

https://api-dumpster-fire.herokuapp.com/

### Running local front end

```
./scripts/frontend_local.sh
```

### Deploy

Will only work for me

```
./scripts/deploy.sh
```