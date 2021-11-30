
# # Text_Title = 'Ejercicio 1'
# # dashtitle = '-' * int((int(78) - len(Text_Title)) / 2)
# # print(dashtitle + Text_Title + dashtitle)
# # # Escribir un programa que pregunte el nombre del usuario en la consola y un número entero e imprima por pantalla en líneas distintas el nombre del usuario tantas veces como el número introducido.

# # nombre = input("Ingrese su nombre: ")
# # numero = int(input("Ingrese un Numero:"))

# # print((nombre+"\n")*numero)

# # 3338344640 numero marcado por Zahara el 24/11/21 a las 9:37
# #-----------------------------------------------------------------------------#


# # Text_Title = 'Ejercicio 2'
# # dashtitle = '-'*int((int(78)-len(Text_Title))/2)
# # print(dashtitle + Text_Title + dashtitle)

# # # Escribir un programa que pregunte el nombre completo del usuario en la consola y después muestre por pantalla el nombre completo del usuario tres veces, una con todas las letras minúsculas, otra con todas las letras mayúsculas y otra solo con la primera letra del nombre y de los apellidos en mayúscula. El usuario puede introducir su nombre combinando mayúsculas y minúsculas como quiera.

# # nombre = input("Ingrese su nombre completo: ")

# # print(nombre.upper())
# # print(nombre.lower())
# # print(nombre.title())

# #-----------------------------------------------------------------------------#


# Text_Title = 'Ejercicio 3'
# dashtitle = '-'*int((int(78)-len(Text_Title))/2)
# print(dashtitle, Text_Title, dashtitle)

# # Escribir un programa que pregunte el nombre del usuario en la consola y después de que el usuario lo introduzca muestre por pantalla <NOMBRE> tiene <n> letras, donde <NOMBRE> es el nombre de usuario en mayúsculas y <n> es el número de letras que tienen el nombre.

# nombre = input("Ingrese su nombre completo: ")

# print(nombre," tiene ",len(nombre), "letras")

#-----------------------------------------------------------------------------#

# Text_Title = 'Ejercicio 4'
# dashtitle = '-'*int((int(78)-len(Text_Title))/2)
# print(dashtitle + Text_Title + dashtitle)

# # Los teléfonos de una empresa tienen el siguiente formato prefijo-número-extension donde el prefijo es el código del país +34, y la extensión tiene dos dígitos (por ejemplo +34-913724710-56). Escribir un programa que pregunte por un número de teléfono con este formato y muestre por pantalla el número de teléfono sin el prefijo y la extensión.

# tel = input("Ingrese su telefono con el formato correspondiente: ")

# print(tel[4:14])
#-----------------------------------------------------------------------------#


# Text_Title = 'Ejercicio 5'
# dashtitle = '-'*int((int(78)-len(Text_Title))/2)
# print(dashtitle + Text_Title + dashtitle)

# # Escribir un programa que pida al usuario que introduzca una frase en la consola y muestre por pantalla la frase invertida.

# frase = input("Ingrese la frase: ")[::-1]

# print(frase)
#-----------------------------------------------------------------------------#


# Text_Title = 'Ejercicio 6'
# dashtitle = '-'*int((int(78)-len(Text_Title))/2)
# print(dashtitle + Text_Title + dashtitle)

# # Escribir un programa que pida al usuario que introduzca una frase en la consola y una vocal, y después muestre por pantalla la misma frase pero con la vocal introducida en mayúscula.
# frase = input("Ingrese la frase: ")
# vocal = input("Ingrese una vocal en minuscula: ")
# print(frase.replace(vocal,vocal.upper()))

#-----------------------------------------------------------------------------#


# Text_Title = 'Ejercicio 7'
# dashtitle = '-'*int((int(78)-len(Text_Title))/2)
# print(dashtitle + Text_Title + dashtitle)

# # Escribir un programa que pregunte el correo electrónico del usuario en la consola y muestre por pantalla otro correo electrónico con el mismo nombre (la parte delante de la arroba @) pero con dominio ceu.es.

# correo = input("Ingrese su correo electronico: ")
# x= (correo[:correo.find('@')]+'@ceu.es')
# print(x)

#-----------------------------------------------------------------------------#


# Text_Title = 'Ejercicio 8'
# dashtitle = '-'*int((int(78)-len(Text_Title))/2)
# print(dashtitle + Text_Title + dashtitle)

# # Escribir un programa que pregunte por consola el precio de un producto en euros con dos decimales y muestre por pantalla el número de euros y el número de céntimos del precio introducido.

# precio = input("Introduce el precio del producto con dos decimales:  ")
# print(precio[:precio.find('.')], 'pesos y', precio[precio.find('.')+1:], 'centavos.')

#-----------------------------------------------------------------------------#


# Text_Title = 'Ejercicio 9'
# dashtitle = '-'*int((int(78)-len(Text_Title))/2)
# print(dashtitle + Text_Title + dashtitle)

# # Ejercicio 9
# # Escribir un programa que pregunte al usuario la fecha de su nacimiento en formato dd/mm/aaaa y muestra por pantalla, el día, el mes y el año. Adaptar el programa anterior para que también funcione cuando el día o el mes se introduzcan con un solo carácter.

# fecha = input("Introduca su fecha de nacimiento en formato dd/mm/aaaa:  ")

# dia=(fecha[:fecha.find("/")])
# mes=(fecha[fecha.find("/")+1:fecha.rfind("/")])
# año=(fecha[fecha.rfind("/")+1:])

# print("Dia: ",dia)
# print("Mes: ",mes)
# print("Año: ",año)

#-----------------------------------------------------------------------------#


# Text_Title = 'Ejercicio 10'
# dashtitle = '-'*int((int(78)-len(Text_Title))/2)
# print(dashtitle + Text_Title + dashtitle)


# # Escribir un programa que pregunte por consola por los productos de una cesta de la compra, separados por comas, y muestre por pantalla cada uno de los productos en una línea distinta.

# productos = input("Introduzca los pŕoductos separados por una coma: ")

# x=('\n'.join(productos.split(',')))

# print(x)

#-----------------------------------------------------------------------------#


Text_Title = 'Ejercicio 11'
dashtitle = '-'*int((int(78)-len(Text_Title))/2)
print(dashtitle + Text_Title + dashtitle)

# Escribir un programa que pregunte el nombre el un producto, su precio y un número de unidades y muestre por pantalla una cadena con el nombre del producto seguido de su precio unitario con 6 dígitos enteros y 2 decimales, el número de unidades con tres dígitos y el coste total con 8 dígitos enteros y 2 decimales.

producto = input("Introduzca el producto: ")
unidades = int(input("Introduzca la cantidad: "))
precio = float(input("Introduzca el precio: "))

print('{producto}: {unidades:3d} unidades x {precio:9.2f}€ = {total:11.2f}€'.format(producto = producto, unidades = unidades, precio = precio, total = unidades * precio))

