# Class 1

## Variables

- Syntax var <name> <type>
- starts with \_ or inital small if Capital than it can be access by other packages.
- " := " short declaration
  1- When you don't know the inital value than use normal declartion which is var
  2- You can only use var declartion at top level (package level)
  3- Use var keyword when the values are grouped
  4- Short declartion is mostly used and recomened by comunities.
- write precise code
- use common conventions

## Args Slice

- []strings , multiple value

## Print
- println prints the value only
- printf prints the type and value 

## Strings
- `` raw stings literal (un processed)

## Pointers
- All pointers based value are intilizied as nil
  - pointers , slices , maps , interface , channels 

## Error Handling
- use nil values for error handlings if a function returns an error

## Fall through
- use in switch case to jump to next case.
- use simple statement like switch i :=; true {}