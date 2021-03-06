# ding

Because Siri won't set a timer for me on macOS.

## Prerequisite for Linux and FreeBSD

### Linux

libasound2-dev is required. On Ubuntu or Debian, run this command:

```sh
apt install libasound2-dev
```

In most cases this command must be run by root user or through `sudo` command.

### FreeBSD

OpenAL is required. Install openal-soft:

```sh
pkg install openal-soft
```

## Installation

1. Install Go via brew or check https://golang.org/ for other methods
    ```sh
    brew install go
    ```

2. Get the package via go get or git clone
    ```
    go get -u github.com/acmujica/ding
    ```
    ```
    git clone https://github.com/acmujica/ding.git
    ```

3. If you cloned you will need to install the binary

    ```
    make install
    ```

## Usage

#### 30 second timer

    ding 30


#### 25 minute timer running in the background

    ding -m 25 &


#### 42 Hour timer

    ding -h 42

#### Slightly louder and longer alarm

    ding -l 10


## Audio Credits

"Hand Bells, E, Single.wav" by InspectorJ (www.jshaw.co.uk) of Freesound.org

Alarm sound source: https://freesound.org/people/bone666138/sounds/198841/
