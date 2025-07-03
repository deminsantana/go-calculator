Cómo se vería en la práctica:
Si guardaras el código anterior en un archivo calculator.go dentro de una carpeta calculator, y luego ejecutaras:

go doc ./calculator: Te mostraría el comentario a nivel de paquete.

go doc ./calculator Add: Te mostraría la documentación completa de la función Add.

go doc ./calculator Divide: Te mostraría la documentación completa de la función Divide, incluyendo la sección de errores y ejemplos.

go doc ./calculator Calculator: Te mostraría la documentación del tipo Calculator y sus métodos.

Y si subes este código a un repositorio público, pkg.go.dev lo indexaría automáticamente y lo mostraría en un formato web interactivo y navegable, haciendo que tu código sea fácil de entender y usar para otros.