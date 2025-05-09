<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: acme/weather/v1/weather.proto

namespace Acme\Weather\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>acme.weather.v1.Location</code>
 */
class Location extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>float latitude = 1 [json_name = "latitude"];</code>
     */
    protected $latitude = 0.0;
    /**
     * Generated from protobuf field <code>float longitude = 2 [json_name = "longitude"];</code>
     */
    protected $longitude = 0.0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type float $latitude
     *     @type float $longitude
     * }
     */
    public function __construct($data = NULL) {
        \Acme\Weather\V1\GPBMetadata\Weather::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>float latitude = 1 [json_name = "latitude"];</code>
     * @return float
     */
    public function getLatitude()
    {
        return $this->latitude;
    }

    /**
     * Generated from protobuf field <code>float latitude = 1 [json_name = "latitude"];</code>
     * @param float $var
     * @return $this
     */
    public function setLatitude($var)
    {
        GPBUtil::checkFloat($var);
        $this->latitude = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>float longitude = 2 [json_name = "longitude"];</code>
     * @return float
     */
    public function getLongitude()
    {
        return $this->longitude;
    }

    /**
     * Generated from protobuf field <code>float longitude = 2 [json_name = "longitude"];</code>
     * @param float $var
     * @return $this
     */
    public function setLongitude($var)
    {
        GPBUtil::checkFloat($var);
        $this->longitude = $var;

        return $this;
    }

}

