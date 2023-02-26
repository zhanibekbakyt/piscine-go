find . -name "*.sh" -printf '%f\n' | sort -nr | sed -e 's/\.sh$//' j
