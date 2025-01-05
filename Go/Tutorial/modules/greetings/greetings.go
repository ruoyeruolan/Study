package greetings

import "fmt"
import "errors"
import "math/rand"

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormate(), name)
    return message, nil
}


func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)

    for _, name := range names {
        message, err := Hello(name)

        if err != nil {return nil, err}

        messages[name] = message
    }
    return messages, nil
}


func randomFormate() string {
    formates := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    return formates[rand.Intn(len(formates))]
}
