# input must be one, two, three, four, five, six, seven, eight, nine, ten. and must not include nothing else other than one-to-ten.
# put those number one-to-ten word into text as input

# scrambled from oneninetentwosixseven
# expecting array similar to 
# [one, nine, ten, two, six, seven]
# OR
# [1, 9, 10, 2, 6, 7]
input = "nnsoitnexniteeeonvswe"

text_number_list = []
number_list = []

while(len(input)>=3):
  if("w" in input):
    input = input.replace("t", "", 1)
    input = input.replace("w", "", 1)
    input = input.replace("o", "", 1)
    text_number_list.append("two")
    number_list.append(2)
   
  elif ("u" in input):
    input = input.replace("f", "", 1)
    input = input.replace("o", "", 1)
    input = input.replace("u", "", 1)
    input = input.replace("r", "", 1)
    text_number_list.append("four")
    number_list.append(4)
  
  elif ("x" in input):
    input = input.replace("s", "", 1)
    input = input.replace("i", "", 1)
    input = input.replace("x", "", 1)
    text_number_list.append("six")
    number_list.append(6)
  
  elif ("g" in input): 
    input = input.replace("e", "", 1)
    input = input.replace("i", "", 1)
    input = input.replace("g", "", 1)
    input = input.replace("h", "", 1)
    input = input.replace("t", "", 1)
    text_number_list.append("eight")
    number_list.append(8)
    
  elif ("z" in input): 
    input = input.replace("z", "", 1)
    input = input.replace("e", "", 1)
    input = input.replace("r", "", 1)
    input = input.replace("o", "", 1)
    text_number_list.append("zero")
    number_list.append(0)
    
  elif ("o" in input): 
    input = input.replace("o", "", 1)
    input = input.replace("n", "", 1)
    input = input.replace("e", "", 1)
    text_number_list.append("one")
    number_list.append(1)
  
  elif ("h" in input): 
    input = input.replace("t", "", 1)
    input = input.replace("h", "", 1)
    input = input.replace("r", "", 1)
    input = input.replace("e", "", 2)
    text_number_list.append("three")
    number_list.append(3)
  
  elif ("f" in input ):
    input = input.replace("f", "", 1)
    input = input.replace("i", "", 1)
    input = input.replace("v", "", 1)
    input = input.replace("e", "", 1)
    text_number_list.append("five")
    number_list.append(5)
  
  elif ("s" in input ):
    input = input.replace("s", "", 1)
    input = input.replace("e", "", 2)
    input = input.replace("v", "", 1)
    input = input.replace("n", "", 1)
    text_number_list.append("seven")
    number_list.append(7)
  
  elif ("nine" in input ):
    input = input.replace("n", "", 2)
    input = input.replace("i", "", 1)
    input = input.replace("e", "", 1)
    text_number_list.append("nine")
    number_list.append(9)
  
  elif ("ten" in input ):
    input = input.replace("t", "", 1)
    input = input.replace("e", "", 1)
    input = input.replace("n", "", 1)
    text_number_list.append("ten")
    number_list.append(10)


print(number_list)
print(text_number_list)
