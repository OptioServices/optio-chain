# Setting up an RPC Node for Optio

This guide will walk you through the process of building the Optio binary and configuring an RPC node for the Optio blockchain.

## 1. Build the Binary

### Step 1: Install Go
```bash
sudo apt-get update && sudo apt install golang-go
```

### Step 2: Clone the Repository
```bash
git clone https://github.com/OptioServices/optio-chain.git optio
cd optio
```

### Step 3: Build the Binary
```bash
env GOOS=linux GOARCH=amd64 go build -o ~/optio/build/optio ~/optio/cmd/optiod/main.go
```

### Step 4: Set Alias for Convenience
```bash
echo 'alias optiod="~/optio/build/optio"' >> ~/.bashrc
source ~/.bashrc
```

## 2. Configure the RPC Node

### Step 1: Initialize the Node
```bash
optiod init optio --chain-id optio --default-denom uOPT
```

### Step 2: Set Minimum Gas Prices
```bash
optiod config set app minimum-gas-prices 1uOPT
```

### Step 3: Get the Genesis File
First, clone the genesis repository:
```bash
cd ..
git clone https://github.com/OptioServices/mainnet-genesis.git
```

Then, copy the genesis file to the config directory:
```bash
cp /root/mainnet-genesis/genesis.json /root/.optio/config/
```

### Step 4: Configure config.toml for RPC
Edit the `config.toml` file located at `/root/.optio/config/config.toml`.

- Replace all instances of `localhost` and `127.0.0.1` with `0.0.0.0`.
- Set `cors_allowed_origins = ["*"]`. Or optionally set your own cors policy

You can use a text editor like `nano` or `vim` to make these changes.

Additionally, you need to set the `seeds` and `persistent_peers` in `config.toml`. The addresses will be provided to you separately. Add them to the fields in the following format:
```
seeds = "node_id1@ip1:26656"
```

### Step 5: Configure app.toml for RPC
Edit the `app.toml` file located at `/root/.optio/config/app.toml`.

- Replace all instances of `localhost` and `127.0.0.1` with `0.0.0.0`.
- Under the `[api]` section, ensure the following are set:
  ```
  [api]
  enable = true
  enabled-unsafe-cors = true // false if you set a cors policy
  ```
- Under the `[api]` section, optionally set your api address:
  ```
  [api]
  address = "tcp://0.0.0.1317"
  ```
- Under the `[grpc]` section, ensure:
  ```
  [grpc]
  enable = true
  ```

## 3. Start the Node

### Step 1: Create Systemd Service File
Create a new file for the systemd service:
```bash
sudo nano /etc/systemd/system/optio.service
```

Paste the following content into the file:
```
[Unit]
Description=Optio Node
After=network-online.target

[Service]
User=root
ExecStart=/usr/local/bin/optiod start
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

Save and close the file.

### Step 2: Copy the Binary to /usr/local/bin
```bash
sudo cp /root/optio/build/optio /usr/local/bin/optiod
```

### Step 3: Enable and Start the Service
```bash
sudo systemctl enable optio.service
sudo systemctl start optio.service
```

### Step 4: Check the Service Status
```bash
sudo systemctl status optio.service
journalctl -xefu optio -o cat
```

This should give you the logs to verify if the node is running correctly.

---

**Note:** Ensure that you have the correct seed addresses provided to you, as they are necessary for your RPC node to connect to the network.