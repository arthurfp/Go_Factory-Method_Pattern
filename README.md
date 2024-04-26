# Factory Method Pattern Demonstration in Go

## Overview
This repository demonstrates the implementation of the Factory Method design pattern in Go. The project showcases how to create objects without specifying the exact class of object that will be created. The primary focus is on manufacturing different types of vehicles like cars, bikes, and trucks, each represented by its own factory.

## Pattern Description
The Factory Method Pattern is a creational design pattern that provides an interface for creating objects in a superclass but allows subclasses to alter the type of objects that will be created. This is useful for managing products that come in multiple variations but share common characteristics. In this project, distinct factories are used for creating different types of vehicles, demonstrating how the Factory Method can be utilized in a real-world application.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), which demonstrates the use of different vehicle factories.
- **pkg/**
    - **factory/**: Includes the concrete implementations for different vehicle factories.
    - **vehicle/**: Contains the interface and implementations for different types of vehicles.

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/arthurfp/Go_Factory-Method_Pattern.git
cd Go_Factory-Method_Pattern
```

### Running the Application
To run the application, execute:
```bash
go run cmd/main.go
```

### Running the Tests
To execute the tests and verify the functionality:
```bash
go test ./...
```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp