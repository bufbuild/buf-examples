<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: acme/weather/v1/weather.proto

namespace Acme\Weather\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>acme.weather.v1.Weather</code>
 */
class Weather extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.acme.weather.v1.Location location = 1 [json_name = "location"];</code>
     */
    protected $location = null;
    /**
     * Generated from protobuf field <code>float temperature = 2 [json_name = "temperature"];</code>
     */
    protected $temperature = 0.0;
    /**
     * Generated from protobuf field <code>float wind_speed = 3 [json_name = "windSpeed"];</code>
     */
    protected $wind_speed = 0.0;
    /**
     * Generated from protobuf field <code>.acme.weather.v1.Condition condition = 4 [json_name = "condition"];</code>
     */
    protected $condition = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Acme\Weather\V1\Location $location
     *     @type float $temperature
     *     @type float $wind_speed
     *     @type int $condition
     * }
     */
    public function __construct($data = NULL) {
        \Acme\Weather\V1\GPBMetadata\Weather::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.acme.weather.v1.Location location = 1 [json_name = "location"];</code>
     * @return \Acme\Weather\V1\Location|null
     */
    public function getLocation()
    {
        return $this->location;
    }

    public function hasLocation()
    {
        return isset($this->location);
    }

    public function clearLocation()
    {
        unset($this->location);
    }

    /**
     * Generated from protobuf field <code>.acme.weather.v1.Location location = 1 [json_name = "location"];</code>
     * @param \Acme\Weather\V1\Location $var
     * @return $this
     */
    public function setLocation($var)
    {
        GPBUtil::checkMessage($var, \Acme\Weather\V1\Location::class);
        $this->location = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>float temperature = 2 [json_name = "temperature"];</code>
     * @return float
     */
    public function getTemperature()
    {
        return $this->temperature;
    }

    /**
     * Generated from protobuf field <code>float temperature = 2 [json_name = "temperature"];</code>
     * @param float $var
     * @return $this
     */
    public function setTemperature($var)
    {
        GPBUtil::checkFloat($var);
        $this->temperature = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>float wind_speed = 3 [json_name = "windSpeed"];</code>
     * @return float
     */
    public function getWindSpeed()
    {
        return $this->wind_speed;
    }

    /**
     * Generated from protobuf field <code>float wind_speed = 3 [json_name = "windSpeed"];</code>
     * @param float $var
     * @return $this
     */
    public function setWindSpeed($var)
    {
        GPBUtil::checkFloat($var);
        $this->wind_speed = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.acme.weather.v1.Condition condition = 4 [json_name = "condition"];</code>
     * @return int
     */
    public function getCondition()
    {
        return $this->condition;
    }

    /**
     * Generated from protobuf field <code>.acme.weather.v1.Condition condition = 4 [json_name = "condition"];</code>
     * @param int $var
     * @return $this
     */
    public function setCondition($var)
    {
        GPBUtil::checkEnum($var, \Acme\Weather\V1\Condition::class);
        $this->condition = $var;

        return $this;
    }

}
