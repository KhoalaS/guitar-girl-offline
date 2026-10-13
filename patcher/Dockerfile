FROM python:3.12-slim-trixie

ENV ANDROID_API=34

RUN apt-get update && apt-get install -y --no-install-recommends curl ca-certificates git unzip jq bash zip wget openjdk-21-jdk

ENV ANDROID_HOME=/opt/android-sdk
ENV PATH=$PATH:$ANDROID_HOME/cmdline-tools/latest/bin
ENV PATH=$PATH:$ANDROID_HOME/build-tools/$ANDROID_API.0.0
ENV ANDROID_SDK_ROOT=$ANDROID_HOME
ENV JAVA_HOME=/usr/lib/jvm/java-21-openjdk-amd64

RUN mkdir -p ${ANDROID_HOME}/cmdline-tools && \
    curl https://dl.google.com/android/repository/commandlinetools-linux-15859902_latest.zip -o /tmp/cmdline-tools.zip && \
    mkdir $ANDROID_HOME/cmdline-tools/latest && \
    unzip /tmp/cmdline-tools.zip && \
    mv cmdline-tools/* $ANDROID_HOME/cmdline-tools/latest && \
    rm /tmp/cmdline-tools.zip && rm -rf cmdline-tools

RUN echo "yes" | sdkmanager --licenses \
    && echo "yes" | sdkmanager "build-tools;$ANDROID_API.0.0"

# create debug keystore
RUN keytool -genkeypair \
      -v \
      -keystore ~/.android/debug.keystore \
      -alias androiddebugkey \
      -keyalg RSA \
      -keysize 2048 \
      -validity 10000 \
      -storepass android \
      -keypass android \
      -dname "CN=Android Debug,O=Android,C=US"

WORKDIR /patcher
COPY app.apkm patch.sh ./
      
RUN git clone https://github.com/jozsefsallai/il2cpp-stringliteral-patcher.git

# patch the apk 
ARG GAME_API_URL="https://localhost"
ENV GAME_API_URL=$GAME_API_URL
RUN bash ./patch.sh

CMD ["tail", "-f"]