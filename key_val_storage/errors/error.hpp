#pragma once

#include <exception>
#include <string>

/// @brief Exception thrown when a key is not found in the key-value store.
// noexcept: Introduced in C++11, this is a strict guarantee to the compiler that this function will never throw an exception of its own. This is critical for error-handling code—if your error-reporting mechanism throws an error while trying to report an error, the C++ program will abruptly terminate (std::terminate). By marking the what() method as noexcept, we ensure that it can be safely called even in situations where exceptions are being handled, without risking further exceptions being thrown. This is a common practice for exception classes in C++.
class KeyNotFoundException : public std::exception {
private:
    std::string message;
public:
    explicit KeyNotFoundException(const std::string& key) {
        message = "Key '" + key + "' not found in the key-value store.";
    }

    const char* what() const noexcept override {
        return message.c_str();
    }
};