#include <iostream>
#include <vector>
#include <unordered_map>
#include "services/IKvManager.hpp"
#include "services/impl/KvManagerImpl.cpp"
using namespace std;

enum class CommandType {
    PUT,
    GET,
    REMOVE,
    EXIT,
    UNKNOWN
};

int main() {
    unordered_map<string, CommandType> commandMap = {
        {"put", CommandType::PUT},
        {"get", CommandType::GET},
        {"remove", CommandType::REMOVE},
        {"exit", CommandType::EXIT}
    };

    IKVManager* kvManager = new KvManagerImpl();
    
    while(true) {
        cout << "Enter command (put/get/remove/exit): ";
        string commandStr;
        if (!(cin >> commandStr)) {
            break; // Exit loop if the input stream fails (e.g., EOF or Ctrl+D)
        }

        CommandType cmd = CommandType::UNKNOWN;
        if (commandMap.find(commandStr) != commandMap.end()) {
            cmd = commandMap[commandStr];
        }

        if (cmd == CommandType::EXIT) {
            break;
        }

        if (cmd == CommandType::UNKNOWN) {
            cout << "Invalid command. Please try again." << endl;
            continue;
        }

        cout << "Enter key: ";
        string key;
        if (!(cin >> key)) break;

        switch (cmd) {
            case CommandType::PUT: {
                cout << "Enter value: ";
                string value;
                if (cin >> value) {
                    kvManager->put(key, value);
                    cout << "Key-value pair ('" << key << "', '" << value << "') added." << endl;
                }
                break;
            }
            case CommandType::GET: {
                try {
                    string value = kvManager->get(key);
                    cout << "Value: " << value << endl;
                } catch (const KeyNotFoundException& e) {
                    cout << e.what() << endl;
                }
                break;
            }
            case CommandType::REMOVE: {
                try {
                    kvManager->remove(key);
                    cout << "Key '" << key << "' removed successfully." << endl;
                } catch (const KeyNotFoundException& e) {
                    cout << e.what() << endl;
                }
                break;
            }
            default:
                break;
        }
    }
    
    cout<<"Exiting the KV Manager."<<endl;
    delete kvManager;
    return 0;
}