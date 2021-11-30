# Crear la parte logica basica de un programa que calcule el el RFC con
# homoclave para México

# REGLAS PARA GENERAR EL RFC
# En más de una ocasión nos encontraremos con algunas ambigüedades al momento
# de extraer las letras de los apellidos o nombres. Es por eso que debemos
# tener en cuenta las siguientes reglas al momento de generar el RFC:

# Si el nombre es compuesto, se considerará el primer nombre excepto cuando sea
# María o José, en cuyo caso se utilizará el segundo nombre.
# Si alguno de los apellidos es compuesto, se considerará la primera palabra
# para formar el RFC.
# Cuando en el nombre o apellidos figuren artículos, preposiciones,
# conjunciones o contracciones no se tomarán en cuenta para la
# generación del RFC.
# Cuando el apellido paterno se componga de una o dos letras entonces las
# letras del RFC se formarán de la siguiente manera: la primera letra del
# apellido paterno seguida de la primera letra del apellido materno y la
# primera y segunda letra del nombre.
# En caso de tener un solo apellido, las letras del RFC serán la primera y
# segunda letra del apellido más la primera y segunda letra del nombre.
# Cuando las cuatro letras del RFC formen una palabra inconveniente, la última
# letra será sustituida por una “X”.
# Otra función que necesitaremos implementar será la que nos ayude a suprimir
# los artículos, preposiciones, conjunciones y contracciones contenidas en el
# nombre o apellidos ya que, como lo establece la regla, dichas palabras no
# serán utilizadas en la generación del RFC.
# El detalle de cada una de estas reglas podrás consultarlo en el documento
# publicado por el IFAI que mencioné al principio del artículo.


Text_Title = 'Calcula RFC Persona Fisica'
dashtitle = '-' * int((int(78) - len(Text_Title)) / 2)
print(dashtitle + Text_Title + dashtitle)

# Declaramos las variables
nomUpper = ''
nombreNoPreps = ''
nom1 = ''
nom2 = ''
paterno = ''
materno = ''

# Solicitamos el nombre completo de la persona

nomCompleto = str(input('Ingrese su nombre completo: '))

# Solicitamos su Fecha de Nacimiento

fechaNacimiento = str(
    input('Ingrese fecha de nacimiento con el formato dd / mm / aaaa: '))

# Damos formato al nombre al reemplazar cualquier vocal acentuada poniendo
# el texto en mayusculas lo cual es de suma importancia porque
# el RFC no admite acentos.


def formatoAlNombre(nomCompleto):
    global nomUpper
    vocalesConAcento = {"Á": "A", "É": "E", "Í": "I", "Ó": "O", "Ú": "U"}
    nomUpper = nomCompleto.upper()

    for letraAcen in vocalesConAcento.keys():
        nomUpper = nomUpper.replace(
            letraAcen, vocalesConAcento.get(letraAcen))
    return nomUpper


# Continuamos dando formato al nombre al reemplazar cualquier vocal acentuada
# poniendo el texto en mayusculas lo cual es de suma importancia porque el RFC
# no admite preposiciones o artuculos como en "del Campo".


def removerPreposiciones():
    global nomUpper
    prepsArtics = (" PARA ", " AND ", " CON ", " DEL ", " LAS ", " LOS ",
                   " MAC ", " POR ", " SUS ", " THE ", " VAN ", " VON ",
                   " AL ", " DE ", " EL ", " EN ", " LA ", " MC ", " MI ",
                   " OF ", " A ", " E ", " Y ")

    for prep in prepsArtics:
        nomUpper = nomUpper.replace(prep, " ")

    return nomUpper

# Continuamos con el formato y removemos los nombres comunes solo si la persona
# tiene un segundo nombre y abreviaciones conforme a las reglas estipuladas


def removerNombresComunes():
    global nomUpper
    nombresComunes = (" MARIA ", " MARIA", "MARIA ", "JOSE ", " JOSE ",
                      " JOSE", " MA. ", " MA ", " J. ", " J ")
    for nombreComun in nombresComunes:
        nomUpper = nomUpper.replace(nombreComun, "")

    return nomUpper

# Dividir el nombre y obtener las letras que componen el RFC y agregar la fecha


def dividirElNombre():
    global nomUpper
    global fechaNacimiento
    nomUpper = nomUpper.split(' ')
    fechaNacimiento = fechaNacimiento.split('/')

    if len(nomUpper) == 4:
        nom1 = nomUpper[0]
        nom2 = nomUpper[1]
        paterno = nomUpper[2]
        materno = nomUpper[3]
        letras1 = paterno[0:2]
        letras2 = materno[0]
        letras3 = nom1[0]

        print(letras1, letras2, letras3)
        print(nom1, nom2, paterno, materno)

    if len(nomUpper) != 4:
        nom1 = nomUpper[0]
        paterno = nomUpper[1]
        materno = nomUpper[2]
        paterno = nomUpper[2]
        materno = nomUpper[3]
        letras1 = paterno[0:1]
        letras2 = materno[0]
        letras3 = nom1[0]
        print(letras1, letras2, letras3)
        print(nom1, nom2, paterno, materno)

    año = fechaNacimiento[2][2:]
    mes = fechaNacimiento[1]
    dia = fechaNacimiento[0]
    print(año, mes, dia)


formatoAlNombre(nomCompleto)
removerPreposiciones()
removerNombresComunes()
dividirElNombre()

#  print(nomUpper)
