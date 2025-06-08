#!/usr/bin/env bash

set -euo pipefail

if [[ "$(uname)" != "Linux" ]]; then
  echo "❌ This script only supports Linux."
  exit 1
fi

GIT_ROOT="$(git rev-parse --show-toplevel 2>/dev/null || echo ".")"
if [[ ! -d "$GIT_ROOT" ]]; then
  echo "❌ Could not determine Git root directory."
  exit 1
fi

# Ensure we are in the Git root directory
cd "$GIT_ROOT"

# Create temporary working directory
TMP_DIR="$(mktemp -d)"
#cleanup() {
#  echo "🧹 Cleaning up..."
#  rm -rf "$TMP_DIR"
#}
#trap cleanup EXIT

echo "📁 Using temporary directory: $TMP_DIR"

echo "🔧 Installing dependencies..."
sudo dnf install -y \
  cmake gcc gcc-c++ make unzip git \
  libX11-devel libXrandr-devel libXcursor-devel libXinerama-devel libXi-devel \
  mesa-libGL-devel mesa-libEGL-devel freetype-devel openal-soft-devel libsndfile-devel \
  curl tar

REPO_DIR="$(pwd)"
cd "$TMP_DIR"

echo "🌱 Cloning SFML 2.6.x from GitHub..."
git clone --depth=1 --branch 2.6.x https://github.com/SFML/SFML.git
cd SFML
mkdir build && cd build
SFML_DIR="$TMP_DIR/SFML"

echo "⚙️ Building SFML from source..."
cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_SHARED_LIBS=OFF
make -j"$(nproc)"
sudo make install

# Return to temp dir
cd "$TMP_DIR"

echo "🌐 Downloading CSFML 2.6.1 source..."
curl -LO https://www.sfml-dev.org/files/CSFML-2.6.1-sources.zip
unzip CSFML-2.6.1-sources.zip
CSFML_DIR="$TMP_DIR/CSFML-2.6.1"

echo "⚙️ Building CSFML from source..."
mkdir -p "$CSFML_DIR/build"
cd "$CSFML_DIR/build"

cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_SHARED_LIBS=OFF
make -j"$(nproc)"
sudo make install

# Copy result to dependencies structure
CSFML_DEST_DIR="$REPO_DIR/dependencies/CSFML_2.6.1"
SFML_DEST_DIR="$REPO_DIR/dependencies/SFML_2.6.1"
mkdir -p "$CSFML_DEST_DIR/include"
mkdir -p "$CSFML_DEST_DIR/lib"
mkdir -p "$SFML_DEST_DIR/include"
mkdir -p "$SFML_DEST_DIR/lib"

# CSFML
cp -r "$CSFML_DIR/include/SFML" "$CSFML_DEST_DIR/include/"
cp -r "$CSFML_DIR/build/lib/"* "$CSFML_DEST_DIR/lib/"

# SFML
cp -r "$SFML_DIR/include/SFML" "$SFML_DEST_DIR/include/"
cp -r "$SFML_DIR/build/lib/"* "$SFML_DEST_DIR/lib/"
#cp /usr/local/lib/libsfml-*.so* "$SFML_DEST_DIR/lib/"

echo ""
echo "✅ SFML + CSFML built and installed"
echo "📦 Files copied to ./dependencies/SFML_2.6.1 and ./dependencies/CSFML_2.6.1"
echo ""
echo "📌 Suggested environment setup:"
echo "export CGO_CFLAGS=\"-I$REPO_DIR/dependencies/CSFML_2.6.1/include -I$REPO_DIR/dependencies/SFML_2.6.1/include\""
echo "export CGO_LDFLAGS=\"-L$REPO_DIR/dependencies/CSFML_2.6.1/lib -L$REPO_DIR/dependencies/SFML_2.6.1/lib -lcsfml-graphics -lcsfml-window -lcsfml-system -lcsfml-network\""

echo ""
echo "(TMP_DIR: $TMP_DIR)"