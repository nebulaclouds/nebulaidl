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

from nebulaadmin.models.core_structured_dataset_metadata import CoreStructuredDatasetMetadata  # noqa: F401,E501


class CoreStructuredDataset(object):
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
        'uri': 'str',
        'metadata': 'CoreStructuredDatasetMetadata'
    }

    attribute_map = {
        'uri': 'uri',
        'metadata': 'metadata'
    }

    def __init__(self, uri=None, metadata=None):  # noqa: E501
        """CoreStructuredDataset - a model defined in Swagger"""  # noqa: E501

        self._uri = None
        self._metadata = None
        self.discriminator = None

        if uri is not None:
            self.uri = uri
        if metadata is not None:
            self.metadata = metadata

    @property
    def uri(self):
        """Gets the uri of this CoreStructuredDataset.  # noqa: E501


        :return: The uri of this CoreStructuredDataset.  # noqa: E501
        :rtype: str
        """
        return self._uri

    @uri.setter
    def uri(self, uri):
        """Sets the uri of this CoreStructuredDataset.


        :param uri: The uri of this CoreStructuredDataset.  # noqa: E501
        :type: str
        """

        self._uri = uri

    @property
    def metadata(self):
        """Gets the metadata of this CoreStructuredDataset.  # noqa: E501


        :return: The metadata of this CoreStructuredDataset.  # noqa: E501
        :rtype: CoreStructuredDatasetMetadata
        """
        return self._metadata

    @metadata.setter
    def metadata(self, metadata):
        """Sets the metadata of this CoreStructuredDataset.


        :param metadata: The metadata of this CoreStructuredDataset.  # noqa: E501
        :type: CoreStructuredDatasetMetadata
        """

        self._metadata = metadata

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
        if issubclass(CoreStructuredDataset, dict):
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
        if not isinstance(other, CoreStructuredDataset):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
