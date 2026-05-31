#include "../IKvManager.hpp"
#include "../../bits/stdc++.h"
#include "../../errors/error.hpp"

using namespace std; 


class KvManagerImpl : public IKVManager {
private:
    unordered_map<string, string> kvStore;
public:
    void put(const string& key, const string& value) override {
        kvStore[key] = value;
    }
    string get(const string& key) override {
        if (kvStore.count(key) == 0) {
            throw KeyNotFoundException(key);
        }  
        return kvStore[key];
    }
    void remove(const string& key) override {
        if (kvStore.count(key) == 0) {
            throw KeyNotFoundException(key);
        }  
        kvStore.erase(key);
    }
};