# Key Value Storage

## Build Instructions

Compile the program with debugging symbols:
```bash
g++ -std=c++11 -g main.cpp -o a.out
```

## Run Instructions

Execute the program:
```bash
./a.out
```

## Debug Instructions

### Setup (One-time)
1. Install CodeLLDB extension in VS Code
2. Ensure your `launch.json` has the correct configuration for CodeLLDB

### Running the Debugger
1. Open VS Code with the project
2. Press `Cmd+Shift+D` to open the Run and Debug view
3. Select the "Debug" configuration from the dropdown
4. Click the green play button to start debugging

### Setting Breakpoints
1. Click on the left margin (gutter) next to the line number where you want to pause
2. A red dot will appear indicating the breakpoint is set

### Debug Controls
- **Continue (F5)**: Resume program execution
- **Step Over (F10)**: Execute the current line and move to the next
- **Step Into (F11)**: Step into function calls
- **Step Out (Shift+F11)**: Step out of the current function

### Inspecting Variables
- Hover over variables to see their current values
- Use the Variables panel on the left to inspect all local variables
- Use the Watch panel to monitor specific expressions

## Cleanup
To remove build artifacts:
```bash
rm -f a.out
rm -rf a.out.dSYM/
```

## Sample execution

```
Enter command (put/get/remove/exit): put
Enter key: sid 
Enter value: raja
Key-value pair ('sid', 'raja') added.
Enter command (put/get/remove/exit): get
Enter key: sid
Value: raja
Enter command (put/get/remove/exit): get
Enter key: raja
Key 'raja' not found in the key-value store.
Enter command (put/get/remove/exit): remove
Enter key: raja
Key 'raja' not found in the key-value store.
Enter command (put/get/remove/exit): remove
Enter key: sid
Key 'sid' removed successfully.
Enter command (put/get/remove/exit): get    
Enter key: sid
Key 'sid' not found in the key-value store.
Enter command (put/get/remove/exit): exit
Exiting the KV Manager.
```