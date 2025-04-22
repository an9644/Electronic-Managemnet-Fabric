# 🔗 Electronic-Managemnet-Fabric
## Hyperledger Fabric Chaincode

This project implements a **role-based access control**  system for managing electronic items in a supply chain using **Hyperledger Fabric**. The smart contract defines operations like creating, reading, and deleting electronic items  with access restricted to specific organizations such as manufacturers.

---

## 🧰 Tech Stack

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Hyperledger Fabric](https://img.shields.io/badge/Hyperledger%20Fabric-101010?style=for-the-badge&logo=hyperledger&logoColor=white)
![CouchDB](https://img.shields.io/badge/CouchDB-DC2626?style=for-the-badge&logo=apachecouchdb&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Minifabric](https://img.shields.io/badge/Minifabric-blue?style=for-the-badge)

---

## 📦 Features

- ✅ Create, Read, and Delete Electronic Items  
- ✅ MSP ID-Based Role Validation  
- ✅ Manufacturer-Only Access to Critical Actions  
- ✅ Uses CouchDB as World State

---

## 🧱 Smart Contract Functions

| Function       | Description                                    | Access Control                     |
|----------------|------------------------------------------------|------------------------------------|
| `CreateItem`   | Adds a new electronic item                     | Only Manufacturer                  |
| `ReadItem`     | Retrieves an item by ID                        | Public                             |
| `DeleteItem`   | Deletes an item by ID                          | Only Manufacturer                  |
| `ItemExists`   | Checks if an item exists in the ledger         | Public                             |

---

## 🧪 Usage with Minifabric

### 🔹 Create an Item

```bash
minifab invoke -n Electronic-Fabric -p '"CreateItem","EL01","WashingMachine","Auto120","Gray","Whirlpool"'
