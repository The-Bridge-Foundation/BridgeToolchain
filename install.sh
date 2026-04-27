#!/bin/sh
echo "Installing Bridge..."

BRIDGE_INSTALL_DIR="$HOME/.bridge"

mkdir -p "$BRIDGE_INSTALL_DIR/bin"
mkdir -p "$BRIDGE_INSTALL_DIR/packages"
mkdir -p "$BRIDGE_INSTALL_DIR/annotations"
cp ./build/bridgec "$BRIDGE_INSTALL_DIR/bin/bridgec"
cp ./build/bridger "$BRIDGE_INSTALL_DIR/bin/bridger"

# Add BRIDGEPATH to shell config files if not already set
BRIDGE_EXPORT="export BRIDGEPATH=\"$BRIDGE_INSTALL_DIR\""
BRIDGE_PATH_EXPORT="export PATH=\"\$BRIDGEPATH/bin:\$PATH\""

add_to_shell_config() {
    local config_file="$1"
    if [ -f "$config_file" ]; then
        if ! grep -q "BRIDGEPATH" "$config_file"; then
            echo "" >> "$config_file"
            echo "# Bridge toolchain" >> "$config_file"
            echo "$BRIDGE_EXPORT" >> "$config_file"
            echo "$BRIDGE_PATH_EXPORT" >> "$config_file"
            echo "Added BRIDGEPATH to $config_file"
        else
            echo "BRIDGEPATH already set in $config_file"
        fi
    fi
}

add_to_shell_config "$HOME/.zshrc"
add_to_shell_config "$HOME/.bashrc"
add_to_shell_config "$HOME/.bash_profile"

echo "Bridge installed to $BRIDGE_INSTALL_DIR"
echo "Restart your shell or run: export BRIDGEPATH=\"$BRIDGE_INSTALL_DIR\""
