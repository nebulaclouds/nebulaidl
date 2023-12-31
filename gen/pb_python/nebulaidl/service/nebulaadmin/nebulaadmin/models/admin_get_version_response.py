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

from nebulaadmin.models.admin_version import AdminVersion  # noqa: F401,E501


class AdminGetVersionResponse(object):
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
        'control_plane_version': 'AdminVersion'
    }

    attribute_map = {
        'control_plane_version': 'control_plane_version'
    }

    def __init__(self, control_plane_version=None):  # noqa: E501
        """AdminGetVersionResponse - a model defined in Swagger"""  # noqa: E501

        self._control_plane_version = None
        self.discriminator = None

        if control_plane_version is not None:
            self.control_plane_version = control_plane_version

    @property
    def control_plane_version(self):
        """Gets the control_plane_version of this AdminGetVersionResponse.  # noqa: E501


        :return: The control_plane_version of this AdminGetVersionResponse.  # noqa: E501
        :rtype: AdminVersion
        """
        return self._control_plane_version

    @control_plane_version.setter
    def control_plane_version(self, control_plane_version):
        """Sets the control_plane_version of this AdminGetVersionResponse.


        :param control_plane_version: The control_plane_version of this AdminGetVersionResponse.  # noqa: E501
        :type: AdminVersion
        """

        self._control_plane_version = control_plane_version

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
        if issubclass(AdminGetVersionResponse, dict):
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
        if not isinstance(other, AdminGetVersionResponse):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
