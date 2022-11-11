package greetings

import (
    "fmt"
    "errors"
    "math/rand"
    "time"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name") 
    }

    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    // map to associate names with messages
    messages := make(map[string]string)
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }

        messages[name] = message
    }
    return messages, nil
}

// execute before all
func init() {
    rand.Seed(time.Now().UnixNano())
}

// return random messages
func randomFormat() string {
    formats := []string{
        "Hi, %v Welcome",
        "Greet to see you, %v",
        "Because you read this %v 🙄",
    }
    
    // Get string array, randomize and return
    return formats[rand.Intn(len(formats))]
}
