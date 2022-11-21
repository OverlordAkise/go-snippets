# Config

I tested and used multiple config file formats in my projects. My summarized, subjective opinions:

## yaml
The perfect, but sadly not inbuilt solution.  
YAML is an asy to read and manage config format.  
Takes only 2 packages (one from gopkg, one from k8s) to parse yaml files easily and quickly.

## json
Very easy to parse thanks to the built-in golang json parser.  
It doesn't support comments and it's a bit weird to manage compared to other formats like xml.

## xml
The easiest to parse and also pretty readable with syntax highlighting.  
Sadly very clustered without colorful highlighting and also way too much text for a simple config file.  

## text / csv
Not as easy to parse as the above and also not really "standardized".  
May be easy to parse but im unsure if it can feature dictionary based configuration easily (with parsing and syntax checking).

