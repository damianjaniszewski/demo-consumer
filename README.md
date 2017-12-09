# demo-consumer

This is a demo-consumer application.

To deploy this application, execute the following commands:

  1. Clone repo:

    ```
    $ git clone https://github.com/damianjaniszewski/demo-consumer
    $ cd demo-consumer
    ```

  2. Install Godep package manager (git required to complete):

    ```
    $ go get github.com/tools/godep

    ```

  3. Create Godep package manager files:

    ```
    $ godep save
    ```

  3. Deploy to HPE Stackato PaaS

    ```
    $ stackato push --stack cflinuxfs2 -n
    ```
