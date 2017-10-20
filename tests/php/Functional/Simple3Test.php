<?php
namespace functional;

class Simple3Test extends FunctionalTestCase
{
    public function testSomeActivitySimulation(): void
    {
        sleep(4);
        $this->assertTrue(
            true
        );
    }

}