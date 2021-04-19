:tocdepth: 2

面板服务 API
^^^^^^^^^^^^

本部分介绍面板服务 API，提供有关如何使用它的指导说明。

.. rest_expand_all::


Login authenticate
""""""""""""""""""

.. rest_method:: POST /api/v1/auth/login

.. rest_parameters:: parameters.yaml

    - username: username
    - password: password

.. rest_parameters:: parameters.yaml

    - user: user
     - userid: userid
     - username: username

.. rest_status_code:: success status.yaml

    - 200

.. rest_status_code:: error status.yaml

    - 400
    - 401
    - 500

logout authenticate
"""""""""""""""""""

.. rest_method:: GET /api/v1/auth/logout

.. rest_status_code:: success status.yaml

    - 200

.. rest_status_code:: error status.yaml

    - 500