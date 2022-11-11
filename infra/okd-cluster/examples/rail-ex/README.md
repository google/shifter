Rails Sample App on OpenShift
============================

This is a quickstart Rails application for OpenShift v3 that you can use as a starting point to develop your own application and deploy it on an [OpenShift](https://github.com/openshift/origin) cluster.

If you'd like to install it, follow [these directions](https://github.com/openshift/rails-ex/blob/master/README.md#installation).  

The steps in this document assume that you have access to an OpenShift deployment that you can deploy applications on.

OpenShift Considerations
------------------------
These are some special considerations you may need to keep in mind when running your application on OpenShift.

###Assets
Your application is set to precompile the assets every time you push to OpenShift.
Any assets you commit to your repo will be preserved alongside those which are generated during the build.

By adding the ```DISABLE_ASSET_COMPILATION=true``` environment variable value to your BuildConfig, you will disable asset compilation upon application deployment.  See the [BuildConfig](http://docs.openshift.org/latest/dev_guide/builds.html#buildconfig-environment) documentation on setting environment variables for builds in OpenShift V3.

###Security
Since these quickstarts are shared code, we had to take special consideration to ensure that security related configuration variables are unique across applications. To accomplish this, we modified some of the configuration files. Now instead of using the same default values, OpenShift can generate these values using the generate from logic defined within the template.

OpenShift stores these generated values in configuration files that only exist for your deployed application and not in your code anywhere. Each of them will be unique so initialize_secret(:a) will differ from initialize_secret(:b) but they will also be consistent, so any time your application uses them (even across reboots), you know they will be the same.

TLDR: OpenShift can generate and expose environment variables to your application automatically. Look at this quickstart for an example.

###Development mode
When you develop your Rails application in OpenShift, you can also enable the 'development' environment by setting the RAILS_ENV environment variable for your deploymentConfiguration, using the `oc` client, like:  

		$ oc env dc/rails-postgresql-example RAILS_ENV=development


If you do so, OpenShift will run your application under 'development' mode. In development mode, your application will:  
*  Show more detailed errors in the browser  
*  Skip static assets (re)compilation  

Development environment can help you debug problems in your application in the same way as you do when developing on your local machine. However, we strongly advise you to not run your application in this mode in production.

###Installation: 
These steps assume your OpenShift deployment has the default set of ImageStreams defined.  Instructions for installing the default ImageStreams are available [here](http://docs.openshift.org/latest/admin_guide/install/first_steps.html).  Instructions for installing the default ImageStreams are available [here](http://docs.openshift.org/latest/admin_guide/install/first_steps.html).  If you are defining the set of ImageStreams now, remember to pass in the proper cluster-admin credentials and to create the ImageStreams in the 'openshift' namespace.

1. Fork a copy of [rails-ex](https://github.com/openshift/rails-ex)
2. Clone your repository to your development machine and cd to the repository directory
3. Add a Ruby application from the rails template:

		$ oc new-app openshift/templates/rails-postgresql.json -p SOURCE_REPOSITORY_URL=https://github.com/< yourusername >/rails-ex 

4. Depending on the state of your system, and whether additional items need to be downloaded, it may take around a minute for your build to be started automatically.  If you do not want to wait, run

		$ oc start-build rails-postgresql-example

5. Once the build is running, watch your build progress  

		$ oc build-logs rails-postgresql-example-1

6. Wait for rails-postgresql-example pods to start up (this can take a few minutes):  

		$ oc get pods -w


	Sample output:  

		NAME                    READY     REASON         RESTARTS   AGE
		postgresql-1-vk6ny                   1/1       Running        0          4m
		rails-postgresql-example-1-build     0/1       ExitCode:0     0          3m
		rails-postgresql-example-1-deploy    1/1       Running        0          34s
		rails-postgresql-example-1-prehook   0/1       ExitCode:0     0          32s



7. Check the IP and port the rails-postgresql-example service is running on:  

		$ oc get svc


	Sample output:  

		NAME             LABELS                              SELECTOR              IP(S)           PORT(S)
		postgresql                 template=rails-postgresql-example   name=postgresql                 172.30.197.40    5432/TCP
		rails-postgresql-example   template=rails-postgresql-example   name=rails-postgresql-example   172.30.205.117   8080/TCP


In this case, the IP for rails-postgresql-example rails-postgresql-example is 172.30.205.117 and it is on port 8080.  
*Note*: you can also get this information from the web console.


###Debugging Unexpected Failures

Review some of the common tips and suggestions [here](https://github.com/openshift/origin/blob/master/docs/debugging-openshift.md).

###Adding Webhooks and Making Code Changes
Since OpenShift V3 does not provide a git repository out of the box, you can configure your github repository to make a webhook call whenever you push your code.

1. From the Web Console homepage, navigate to your project
2. Click on Browse > Builds
3. Click the link with your BuildConfig name
4. Click the Configuration tab
5. Click the "Copy to clipboard" icon to the right of the "GitHub webhook URL" field
6. Navigate to your repository on GitHub and click on repository settings > webhooks > Add webhook
7. Paste your webhook URL provided by OpenShift
8. Leave the defaults for the remaining fields - That's it!
9. After you save your webhook, if you refresh your settings page you can see the status of the ping that Github sent to OpenShift to verify it can reach the server.  

#### Enabling the Blog example
In order to access the example blog application, you have to remove the
`public/index.html` which serves as the welcome page and rebuild the application.
Another option is to make a request directly to `/articles` which will give you access to the blog.

The username/pw used for authentication in this application are openshift/secret.

### Hot Deploy

In order to dynamically pick up changes made in your application source code, you need to set the `RAILS_ENV=development` parameter to the [oc new-app](https://docs.openshift.org/latest/cli_reference/basic_cli_operations.html#basic-cli-operations) command, while performing the [installation steps](https://github.com/openshift/rails-ex#installation) described in this README.

	$ oc new-app openshift/templates/rails-postgresql.json -p RAILS_ENV=development

To change your source code in the running container you need to [oc rsh](https://docs.openshift.org/latest/cli_reference/basic_cli_operations.html#troubleshooting-and-debugging-cli-operations) into it.

	$ oc rsh <POD_ID>

After you [oc rsh](https://docs.openshift.org/latest/cli_reference/basic_cli_operations.html#troubleshooting-and-debugging-cli-operations) into the running container, your current directory is set to `/opt/app-root/src`, where the source code is located.

To set your application back to the `production` environment you need to remove `RAILS_ENV` environment variable:

	$ oc env dc/rails-postgresql-example RAILS_ENV-

A redeploy will happen automatically due to the `ConfigChange` trigger.

**NOTICE: If the `ConfigChange`  trigger is not set, you need to run the redeploy manually:**
 
	$ oc deploy rails-postgresql-example --latest

###License
This code is dedicated to the public domain to the maximum extent permitted by applicable law, pursuant to [CC0](http://creativecommons.org/publicdomain/zero/1.0/).
