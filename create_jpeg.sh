while read -r; do
	sips -s format jpeg "$REPLY/cover.png" --out "${REPLY}/cover.jpg"
done