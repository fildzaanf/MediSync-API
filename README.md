# MediSync : Medication Scheduling and Tracking

## 📝 Project Overview
MediSync is a digital health platform that provides services for scheduling medication use times, managing medication stock, and medication information via MediChat integrated with Artificial Intelligence 
 
## 🎯Problem Statement & Solution
#### Problem Statement
Many people struggle to manage their medication routines due to forgotten schedules, untracked medicine stock, and limited access to reliable medication information. These issues can lead to missed doses, improper usage, and reduced treatment effectiveness.

#### Solution
MediSync provides an integrated digital platform for medication scheduling, stock tracking, and information management.
The platform allows users to:

- Schedule medication intake with automated reminders
- Track medicine stock to prevent shortages
- Store and access medication information
- Receive scheduled email notifications
- Interact with an AI-powered chatbot (MediChat) for medication guidance

By combining scheduling automation, background jobs, email notifications, and AI-powered assistance, MediSync helps users maintain consistent medication routines, improve adherence, and manage their medications more effectively in a centralized and user-friendly system.

## 🚀 Tools and Technologies 
- Go Programming Language
- Echo Go Framework
- GORM for Object Relational Mapping
- MySQL for Relational Database
- JSON Web Token for Authentication
- Docker for Containerization
- CI/CD Pipeline using GitHub Actions
- Amazon Web Services (AWS)
  - Amazon Elastic Compute Cloud (EC2)
  - Amazon Relational Database Service (RDS)
- Large Language Model (LLM) with Generative Pre-trained Transformer (GPT)
- ChatBot with Prompt Engineering
- Go Cron Jobs with SMTP and Go Mail to scheduled email reminders
- Go Testify for Unit Test

## 🏛️ System Design and Architecture
- Domain-Driven Design (DDD)
- Model View Controller (MVC) Architecture
- REST API

## ✨ Features
| Feature                   | Description                                                                                                                                                          |
| ------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| User Management           | Enables users to register, log in, and manage their profile information securely using JWT-based authentication                                                      |
| Medication Management     | Allows users to add, update, view, and delete medicines, including storing detailed medication information such as name, dosage, and category                        |
| Schedule Management       | Provides functionality to create and manage medication intake schedules, helping users take medicines at the correct time consistently                               |
| Medical ID Management     | Stores and manages user medical identity data to associate medications, schedules, and health-related information with specific users                                |
| Automated Email Reminders | Sends scheduled email notifications using background jobs (Go Cron) to remind users to take their medications on time                                                |
| MediChat (AI Chatbot)     | AI-powered chatbot integrated with a GPT-based model that helps users obtain medication-related information and basic guidance through natural language interaction  |
| Category Management       | Organizes medicines into categories to improve data structuring, filtering, and easier management                                                                    |


## 📡 API Endpoints

#### Users
| Method | Endpoint        | Description       |
|--------|-----------------|-------------------|
| POST   | /users/register | Register  user   |
| POST   | /users/login    | Login user       |
| GET    | /users          | Get all users    |
| GET    | /users/:id      | Get user by ID   |
| PUT    | /users/:id      | Update user      |
| DELETE | /users/:id      | Delete user      |

#### Medical IDs
| Method | Endpoint        | Description             |
|--------|-----------------|-------------------------|
| POST   | /medicalids/:id | Create medical ID       |
| GET    | /medicalids     | Get all medical IDs     |
| GET    | /medicalids/:id | Get medical ID by ID    |
| PUT    | /medicalids/:id | Update medical ID       |
| DELETE | /medicalids/:id | Delete medical ID       |

#### Medicines
| Method | Endpoint        | Description           |
|--------|-----------------|-----------------------|
| POST   | /medicines/:id  | Create medicine       |
| GET    | /medicines      | Get all medicines     |
| GET    | /medicines/:id  | Get medicine by ID    |
| PUT    | /medicines/:id  | Update medicine       |
| DELETE | /medicines/:id  | Delete medicine       |

#### Schedules
| Method | Endpoint        | Description           |
|--------|-----------------|-----------------------|
| POST   | /schedules/:id  | Create schedule       |
| GET    | /schedules      | Get all schedules     |
| GET    | /schedules/:id  | Get schedule by ID    |
| PUT    | /schedules/:id  | Update schedule       |
| DELETE | /schedules/:id  | Delete schedule       |

#### Categories
| Method | Endpoint        | Description           |
|--------|-----------------|-----------------------|
| POST   | /categories     | Create category       |
| GET    | /categories     | Get all categories    |
| GET    | /categories/:id | Get category by ID    |
| PUT    | /categories/:id | Update category       |
| DELETE | /categories/:id | Delete category       |

#### MediChats
| Method | Endpoint   | Description               |
|--------|------------|---------------------------|
| POST   | /medichats | Create MediChat message   |

