FROM nginx:1.29-alpine

RUN apk add --no-cache bash curl gettext

# Template configs
COPY ./deployment/docker/nginx/templates /etc/nginx/templates/

# Healthcheck script
COPY ./deployment/docker/nginx/healthcheck.sh /healthcheck.sh
RUN chmod +x /healthcheck.sh

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD /healthcheck.sh

EXPOSE 80