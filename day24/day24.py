from z3 import *

var_counts = { 'w': 0, 'x': 0, 'y': 0, 'z': 0, "input": 0 }

# Read the input into a list of instructions
instructions = []
with open("input.txt", "r") as f:
    for line in f.readlines():
        parts = line.strip().split(' ')
        
        op = parts[0]
        
        if op == 'inp':
            
            src = int(var_counts["input"])
            var_counts["input"] += 1

            var_counts[parts[1]] += 1
            dest = parts[1] + str(var_counts[parts[1]])

            instructions.append({
                "op": op,
                "src": src,
                "dest": dest,
            })

        else:
            src1 = parts[1] + str(var_counts[parts[1]])

            if parts[2] in var_counts:
                src2 = parts[2] + str(var_counts[parts[2]])
            else:
                src2 = int(parts[2])

            var_counts[parts[1]] += 1
            dest = parts[1] + str(var_counts[parts[1]])

            instructions.append({
                "op": op,
                "src1": src1,
                "src2": src2,
                "dest": dest,
            })

# Add input variables and constraints
inputs = [Int(f"input{i}") for i in range(14)]
input_constraint_lower = And([i > 0 for i in inputs])
input_constraint_upper = And([i < 10 for i in inputs])
input_constraint = And(input_constraint_lower, input_constraint_upper)

# Add initial register values
variables = {}
initial_registers_constraints = []
for r in ['w', 'x', 'y', 'z']:
    variable = Int(f'{r}0')
    variables[f'{r}0'] = variable
    initial_registers_constraints.append(variable == 0)

# Add constraints from the program instructions
program_constraints = []
for instruction in instructions:

    # Create a new variable to represent the destination register
    variable = Int(instruction["dest"])
    variables[instruction["dest"]] = variable

    if instruction['op'] == "inp":
        constraint = variable == inputs[instruction['src']]
        program_constraints.append(constraint)
    elif instruction['op'] == "mul":
        src1 = variables[instruction["src1"]]
        src2 = variables.get(instruction["src2"], instruction["src2"])
        program_constraints.append(variable == src1 * src2)
    elif instruction['op'] == "add":
        src1 = variables[instruction["src1"]]
        src2 = variables.get(instruction["src2"], instruction["src2"])
        program_constraints.append(variable == src1 + src2)
    elif instruction['op'] == "div":
        src1 = variables[instruction["src1"]]
        src2 = variables.get(instruction["src2"], instruction["src2"])
        program_constraints.append(variable == src1 / src2)
    elif instruction['op'] == "mod":
        src1 = variables[instruction["src1"]]
        src2 = variables.get(instruction["src2"], instruction["src2"])
        program_constraints.append(variable == src1 % src2)
    elif instruction['op'] == "eql":
        src1 = variables[instruction["src1"]]
        src2 = variables.get(instruction["src2"], instruction["src2"])
        program_constraints.append(Implies(src1 == src2, variable == 0))
        program_constraints.append(Implies(src1 != src2, variable == 1))

# Add final constraint (z == 0)
final_register = variables[f"z{var_counts['z']}"]
final_register_constraint = final_register == 0

# Solve for part 1
solver = Solver()
solver.add(And(program_constraints))
solver.add(And(initial_registers_constraints))
solver.add(final_register_constraint)
solver.add(input_constraint)
solver.push()

digits = []
for digit in range(14):
    for i in range(9, -1, -1):
        solver.push()
        solver.add(inputs[digit] == i)
        digits.append(i)

        if solver.check() == sat:
            break
        else:
            solver.pop()
            digits.pop()

print('Day 24 - Part 1: ', ''.join(map(str, digits)))

# Solve part 2
solver = Solver()
solver.add(And(program_constraints))
solver.add(And(initial_registers_constraints))
solver.add(final_register_constraint)
solver.add(input_constraint)
solver.push()

digits = []
for digit in range(14):
    for i in range(1, 10):
        solver.push()
        solver.add(inputs[digit] == i)
        digits.append(i)

        if solver.check() == sat:
            break
        else:
            solver.pop()
            digits.pop()

print('Day 24 - Part 2: ', ''.join(map(str, digits)))