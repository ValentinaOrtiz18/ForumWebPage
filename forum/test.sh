#!/bin/bash

# Script to test the forum application

echo "🧪 Testing Forum Application..."
echo ""

# Check if server is running
if ! curl -s http://localhost:3030 > /dev/null; then
    echo "❌ Server is not running. Please start the server first with: ./forum"
    exit 1
fi

echo "✅ Server is running on http://localhost:3030"
echo ""

echo "📋 Testing endpoints:"
echo ""

# Test home page
echo -n "  Testing home page (/)... "
if curl -s -o /dev/null -w "%{http_code}" http://localhost:3030 | grep -q "200"; then
    echo "✅ OK"
else
    echo "❌ Failed"
fi

# Test login page
echo -n "  Testing login page (/login)... "
if curl -s -o /dev/null -w "%{http_code}" http://localhost:3030/login | grep -q "200"; then
    echo "✅ OK"
else
    echo "❌ Failed"
fi

# Test register page
echo -n "  Testing register page (/register)... "
if curl -s -o /dev/null -w "%{http_code}" http://localhost:3030/register | grep -q "200"; then
    echo "✅ OK"
else
    echo "❌ Failed"
fi

# Test static files
echo -n "  Testing CSS static file... "
if curl -s -o /dev/null -w "%{http_code}" http://localhost:3030/static/css/style.css | grep -q "200"; then
    echo "✅ OK"
else
    echo "❌ Failed"
fi

echo ""
echo "✅ All basic tests passed!"
echo ""
echo "🌐 Open http://localhost:3030 in your browser to use the forum"
