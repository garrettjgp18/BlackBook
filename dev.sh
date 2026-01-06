#!/bin/bash

# BlackBook Development Helper Script

set -e

echo "BlackBook Development Helper"
echo "============================"
echo ""

# Check if wails is installed
if ! command -v wails &> /dev/null; then
    echo "❌ Wails CLI not found. Installing..."
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    export PATH=$PATH:$(go env GOPATH)/bin
fi

# Function to check Ollama
check_ollama() {
    echo "🔍 Checking Ollama status..."
    if curl -s http://localhost:11434/api/tags > /dev/null 2>&1; then
        echo "✅ Ollama is running"
        echo "   Available models:"
        curl -s http://localhost:11434/api/tags | grep -o '"name":"[^"]*"' | cut -d'"' -f4 || echo "   (No models found)"
    else
        echo "⚠️  Ollama is not running or not installed"
        echo "   AI features will not be available"
        echo "   Visit https://ollama.ai to install"
    fi
    echo ""
}

# Function to install dependencies
install_deps() {
    echo "📦 Installing dependencies..."
    
    echo "  - Go modules..."
    go mod download
    
    echo "  - Frontend packages..."
    cd frontend
    npm install
    cd ..
    
    echo "✅ Dependencies installed"
    echo ""
}

# Function to build frontend
build_frontend() {
    echo "🔨 Building frontend..."
    cd frontend
    npm run build
    cd ..
    echo "✅ Frontend built"
    echo ""
}

# Function to run development server
run_dev() {
    echo "🚀 Starting development server..."
    export PATH=$PATH:$(go env GOPATH)/bin
    wails dev
}

# Function to build production
build_prod() {
    echo "🏗️  Building for production..."
    export PATH=$PATH:$(go env GOPATH)/bin
    wails build
    echo "✅ Build complete! Binary is in build/bin/"
    echo ""
}

# Main menu
case "${1:-}" in
    "check")
        check_ollama
        ;;
    "install")
        install_deps
        ;;
    "build-frontend")
        build_frontend
        ;;
    "dev")
        check_ollama
        run_dev
        ;;
    "build")
        install_deps
        build_frontend
        build_prod
        ;;
    *)
        echo "Usage: $0 [command]"
        echo ""
        echo "Commands:"
        echo "  check          - Check Ollama status"
        echo "  install        - Install dependencies"
        echo "  build-frontend - Build frontend only"
        echo "  dev            - Run development server"
        echo "  build          - Build production binary"
        echo ""
        echo "Quick start:"
        echo "  ./dev.sh install   # Install dependencies"
        echo "  ./dev.sh dev       # Run in development mode"
        ;;
esac
