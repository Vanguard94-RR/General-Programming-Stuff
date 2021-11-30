import turtle

colores = ['Yellow', 'Green', 'Red', 'Blue', 'Orange', 'White', 'Cyan',
           'Magenta']

t = turtle.Turtle()
screen = turtle.Screen()
screen.bgcolor('black')
t.speed(5000)

for i in range(120):
    t.color(colores[i % 8])
    t.forward(i * 6)
    t.left(150)
    t.width(4)
