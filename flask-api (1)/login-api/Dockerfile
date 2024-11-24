# Gunakan Node.js versi 20 sebagai base image
FROM node:20

# Set working directory di dalam container
WORKDIR /usr/src/app

# Salin file package.json dan package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Salin semua file proyek ke dalam container
COPY . .

# Ekspos port aplikasi (3000 untuk API)
EXPOSE 3000

# Jalankan server saat container mulai
CMD ["node", "index.js"]