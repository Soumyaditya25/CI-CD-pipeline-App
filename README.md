Thought for a couple of seconds


## About

This project demonstrates a CI/CD pipeline for a simple web application. The app exposes a login API that returns user details in JSON format. Whenever a request is received, it responds with a name and corresponding data.

### Tools Used

* **Docker**: Containerize the application
* **Minikube**: Run a local Kubernetes cluster
* **AWS EC2**: Host a Jenkins server
* **Jenkins**: Manage the continuous integration workflow
* **Argo CD**: Implement continuous deployment via GitOps principles

---

## Workflow Overview

![CI/CD Workflow](https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/f77ce19f-7030-49e5-9402-5921493d478b)

---

## Step 1: Build & Run the Docker Image

1. **Build** the image locally:

   ```bash
   docker build -t soumyaditya/pipeline .
   ```
2. **Run** the container:

   ```bash
   docker run -it -p 3001:3001 soumyaditya/pipeline
   ```

   The app will be accessible on port **3001**.

---

## Step 2: Push to Docker Hub

1. **Authenticate** (if needed):

   ```bash
   docker login
   ```
2. **Push** the image:

   ```bash
   docker push soumyaditya/pipeline
   ```

---

## Step 3: Provision Jenkins on AWS EC2

1. **Launch** an EC2 instance and allow inbound traffic from Anywhere (port 8080). <img width="1325" alt="EC2 Security Group" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/c440ae5a-78ff-492f-8a8e-c4c07b02598a">
2. **Install** Jenkins on the instance (follow the [official docs](https://www.jenkins.io/doc/book/installing/linux/#debianubuntu)):

   ```bash
   sudo apt update
   sudo apt install openjdk-11-jdk -y
   sudo apt install jenkins -y
   systemctl start jenkins
   systemctl status jenkins
   ```
3. **Unlock** Jenkins:

   * Open `http://<EC2_IP>:8080` in your browser.
   * Retrieve the initial admin password:

     ```bash
     sudo cat /var/lib/jenkins/secrets/initialAdminPassword
     ```
   * Install suggested plugins.

---

## Step 4: Configure Jenkins Credentials & Pipeline

1. **Add Credentials**

   * **GitHub**: Secret text (GitHub token)
   * **Docker Hub**: Username & password
2. **Create** a new **Pipeline** project and save.

---

## Step 5: Enable Docker Builds in Jenkins

1. **Install Docker** on the EC2 instance:

   ```bash
   sudo apt-get install docker.io -y
   sudo systemctl start docker
   systemctl status docker
   ```
2. **Grant Docker Access** to Jenkins and Ubuntu users:

   ```bash
   sudo usermod -aG docker jenkins
   sudo usermod -aG docker ubuntu
   sudo chmod 777 /var/run/docker.sock
   systemctl restart jenkins
   ```
3. **Install** the **Docker Pipeline** plugin in Jenkins.
4. **Add** a **GitHub webhook** pointing to:

   ```
   http://<EC2_IP>:8080/github-webhook/
   ```
5. **Trigger** the first build by clicking **Build Now**: <img width="1116" alt="Jenkins Build" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/1dca6396-2fe2-4c74-b60d-33f87dc74279">

> Any push to the `main` branch will now automatically trigger a Jenkins build.

---

## Step 6: Install & Access Argo CD on Minikube

1. **Follow** the [Argo CD Getting Started guide](https://argo-cd.readthedocs.io/en/stable/getting_started/) to deploy in Minikube.
2. **Access** Argo CD at `http://localhost:8080`. <img width="1380" alt="Argo CD UI" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/dfad6712-770c-4a3e-bed1-9622a072e0be">
3. **Log in** using the default credentials per the docs.

---

## Step 7: Deploy the App with Argo CD

1. **Create** a new application:

   * Click **New App**
   * Connect your GitHub repository
   * Select branch **main**
   * Set **Path** to `/`
   * Click **Create**
2. **Verify** deployment:

   ```bash
   kubectl get pods
   ```

   <img width="1017" alt="Kubernetes Pods" src="https://github.com/prady0t/CI-CD-pipeline-app/assets/99216956/608d7e1f-e5c6-4653-bf3f-0ca6fbdf761c">  

> Argo CD will automatically sync and manage your Kubernetes resources as defined in your manifests.

---

With these steps, both the CI (via Jenkins) and CD (via Argo CD) workflows are fully operational.
