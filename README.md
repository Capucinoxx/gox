# gox
Interpréteur regex

## Regex fréquemments utilisé

### Fonction génératrices sur les nombres

| code              | regex                                             |
| ----------------- | ------------------------------------------------- |
| (n)               | `/[-]?[0-9]+[,.]?[0-9]*([\/][0-9]+[,.]?[0-9]*)*/` |
| (n whole)         | `/^\d+/`                                          |
| (n decimal        | `/^\d*\.\d+/`                                     |
| (n whole decimal) | `/^\d*(\.\d+)?/`                                  |


## TODO
- [ ] lexer
- [ ] parser
- [ ] fonctions génératrices des nombres
- [ ] attributs `space` disponible dans chaque fonction
- [ ] composition de fonction