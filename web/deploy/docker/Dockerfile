# 使用 Node.js 18.12.0 版本的官方镜像作为基础镜像
# FROM  registry.cn-chengdu.aliyuncs.com/nexwind_backend/node_18_12_0_alpine:latest AS build-stage
# FROM  registry.cn-chengdu.aliyuncs.com/nexwind_backend/node_18_12_0:latest AS build-stage
# 设置工作目录
# WORKDIR /app

# 将 package.json 和 package-lock.json 复制到容器中
# COPY package*.json ./
# RUN npm config set registry https://registry.npmmirror.com

# 安装生产依赖
# RUN npm ci

# 将源代码复制到容器中
# COPY ../../ .

# 安装开发依赖
# RUN npm install

# 构建 Vite 应用程序
# RUN npm run build

# 使用 Nginx 镜像来服务静态文件
FROM registry.cn-chengdu.aliyuncs.com/nexwind_backend/nginx-alpine:latest AS production-stage

# 复制 nginx 配置文件
COPY ../../deploy/nginx/nginx.conf /etc/nginx/nginx.conf

# 将构建好的文件复制到 Nginx 的静态文件目录
# COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY ../../dist /usr/share/nginx/html

# 暴露端口
EXPOSE 8881

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]