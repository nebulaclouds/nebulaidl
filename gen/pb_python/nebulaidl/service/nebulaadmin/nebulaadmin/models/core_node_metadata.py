# coding: utf-8

"""
    nebulaidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from nebulaadmin.models.core_retry_strategy import CoreRetryStrategy  # noqa: F401,E501


class CoreNodeMetadata(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'name': 'str',
        'timeout': 'str',
        'retries': 'CoreRetryStrategy',
        'interruptible': 'bool'
    }

    attribute_map = {
        'name': 'name',
        'timeout': 'timeout',
        'retries': 'retries',
        'interruptible': 'interruptible'
    }

    def __init__(self, name=None, timeout=None, retries=None, interruptible=None):  # noqa: E501
        """CoreNodeMetadata - a model defined in Swagger"""  # noqa: E501

        self._name = None
        self._timeout = None
        self._retries = None
        self._interruptible = None
        self.discriminator = None

        if name is not None:
            self.name = name
        if timeout is not None:
            self.timeout = timeout
        if retries is not None:
            self.retries = retries
        if interruptible is not None:
            self.interruptible = interruptible

    @property
    def name(self):
        """Gets the name of this CoreNodeMetadata.  # noqa: E501


        :return: The name of this CoreNodeMetadata.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this CoreNodeMetadata.


        :param name: The name of this CoreNodeMetadata.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def timeout(self):
        """Gets the timeout of this CoreNodeMetadata.  # noqa: E501

        The overall timeout of a task.  # noqa: E501

        :return: The timeout of this CoreNodeMetadata.  # noqa: E501
        :rtype: str
        """
        return self._timeout

    @timeout.setter
    def timeout(self, timeout):
        """Sets the timeout of this CoreNodeMetadata.

        The overall timeout of a task.  # noqa: E501

        :param timeout: The timeout of this CoreNodeMetadata.  # noqa: E501
        :type: str
        """

        self._timeout = timeout

    @property
    def retries(self):
        """Gets the retries of this CoreNodeMetadata.  # noqa: E501

        Number of retries per task.  # noqa: E501

        :return: The retries of this CoreNodeMetadata.  # noqa: E501
        :rtype: CoreRetryStrategy
        """
        return self._retries

    @retries.setter
    def retries(self, retries):
        """Sets the retries of this CoreNodeMetadata.

        Number of retries per task.  # noqa: E501

        :param retries: The retries of this CoreNodeMetadata.  # noqa: E501
        :type: CoreRetryStrategy
        """

        self._retries = retries

    @property
    def interruptible(self):
        """Gets the interruptible of this CoreNodeMetadata.  # noqa: E501


        :return: The interruptible of this CoreNodeMetadata.  # noqa: E501
        :rtype: bool
        """
        return self._interruptible

    @interruptible.setter
    def interruptible(self, interruptible):
        """Sets the interruptible of this CoreNodeMetadata.


        :param interruptible: The interruptible of this CoreNodeMetadata.  # noqa: E501
        :type: bool
        """

        self._interruptible = interruptible

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(CoreNodeMetadata, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, CoreNodeMetadata):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
